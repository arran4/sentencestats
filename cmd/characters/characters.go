package main

import (
	"flag"
	"fmt"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"unicode"
)

var (
	outputFile = flag.String("o", "out.png", "Output PNG file name")
)

func main() {
	flag.Parse()
	b, _ := ioutil.ReadAll(os.Stdin)
	s := string(b)
	type Sentence struct {
		Hist     [26]float64
		Sentence string
		Count    int
	}
	sentences := []Sentence{{}}
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			c := unicode.ToLower(r)
			sentences[len(sentences)-1].Hist[c-'a'] += 1
			sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			sentences[len(sentences)-1].Count++
		} else {
			switch r {
			case '.':
				sentences = append(sentences, Sentence{})
			case '\r', '\n':
			case '\t':
				sentences[len(sentences)-1].Sentence += fmt.Sprintf(" ")
			default:
				sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			}
		}
	}
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

	f, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	i := image.NewRGBA(image.Rect(0, 0, 900, 416))
	igr := imgg.AddTo(i, 0, 0, 900, 416, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(i, i.Bounds(), bg, image.Point{}, draw.Src)
	c.Plot(igr)
	png.Encode(f, i)

}
