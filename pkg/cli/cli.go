// Package cli provides CLI subcommands for sentencestats.
package cli

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"

	"github.com/arran4/sentencestats/pkg/analyze"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
)

// Characters is a subcommand `sentencestats characters`
//
// Flags:
//
//	output: -o --output (default: "out.png") Output PNG file name
func Characters(output string) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read from stdin: %v", err)
	}
	s := string(b)

	sentences := analyze.AnalyzeCharacters(s)

	c := chart.BarChart{Title: "Character frequency", Stacked: true}
	x := []float64{}
	c.XRange.Category = []string{}
	for i := 'a'; i <= 'z'; i++ {
		x = append(x, float64(i-'a'))
		c.XRange.Category = append(c.XRange.Category, fmt.Sprintf("%c", i))
	}
	c.XRange.Label = "Character"
	c.YRange.Label = "Frequency #"
	c.YRange.MinMode.Fixed = true
	for i, s := range sentences {
		if s.Count == 0 {
			continue
		}
		co := chart.Style{Symbol: '#', LineColor: color.NRGBA{0x00, 0x00, 0xff, 0xff}, LineWidth: 3, FillColor: color.NRGBA{byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)), 0xff}}
		c.AddDataPair(fmt.Sprintf("Sentence %d: %s", i, s.Sentence), x, s.Hist[:], co)
	}

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Failed to close output file: %v", err)
		}
	}()

	i := image.NewRGBA(image.Rect(0, 0, 900, 416))
	igr := imgg.AddTo(i, 0, 0, 900, 416, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(i, i.Bounds(), bg, image.Point{}, draw.Src)
	c.Plot(igr)
	if err := png.Encode(f, i); err != nil {
		log.Printf("Failed to encode png: %v", err)
	}
}

// CharacterPairs is a subcommand `sentencestats character-pairs`
//
// Flags:
//
//	output: -o --output (default: "out.png") Output PNG file name
func CharacterPairs(output string) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read from stdin: %v", err)
	}
	s := string(b)

	sentences, allPairs := analyze.AnalyzePairs(s)

	c := chart.BarChart{Title: "Character Pair frequency", Stacked: true}
	x := []float64{}
	c.XRange.Label = "Character Pair"
	c.YRange.Label = "Frequency #"
	c.YRange.MinMode.Fixed = true
	ct := 0.0
	order := []string{}
	for p := range allPairs {
		order = append(order, p)
		x = append(x, ct)
		log.Printf("all pair: %s", p)
		ct++
	}
	sort.Strings(order)
	log.Printf("%#v", allPairs)
	c.XRange.Category = order
	for i, s := range sentences {
		if s.Count == 0 {
			continue
		}
		co := chart.Style{Symbol: '#', LineColor: color.NRGBA{0x00, 0x00, 0xff, 0xff}, LineWidth: 3, FillColor: color.NRGBA{byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)), 0xff}}
		vs := []float64{}
		for _, p := range order {
			v := s.Pairs[p]
			vs = append(vs, v)
		}
		log.Printf("%#v", vs)
		log.Printf("%#v", order)
		c.AddDataPair(fmt.Sprintf("Sentence %d: %s", i, s.Sentence), x, vs, co)
	}

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Failed to close output file: %v", err)
		}
	}()

	const width = 1600
	height := 416
	i := image.NewRGBA(image.Rect(0, 0, width, height))
	igr := imgg.AddTo(i, 0, 0, width, height, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(i, i.Bounds(), bg, image.Point{}, draw.Src)
	c.Plot(igr)
	if err := png.Encode(f, i); err != nil {
		log.Printf("Failed to encode png: %v", err)
	}
}
