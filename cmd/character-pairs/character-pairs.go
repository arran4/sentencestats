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
	"log"
	"math/rand"
	"os"
	"sort"
	"unicode"
)

var (
	outputFile = flag.String("o", "out.png", "Output PNG file name")
)

func main() {
	flag.Parse()
	b, _ := ioutil.ReadAll(os.Stdin)
	s := string(b)
	allPairs := map[string]float64{}
	stackPairs := map[string]float64{}
	type Sentence struct {
		Pairs    map[string]float64
		Sentence string
		Count    int
	}
	sentences := []Sentence{{
		Pairs:    map[string]float64{},
		Sentence: "",
		Count:    0,
	}}
	var prev rune = 0
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			c := unicode.ToLower(r)
			if prev > 0 {
				s := string([]rune{c, prev})
				if prev > r {
					s = string([]rune{prev, c})
				}
				allPairs[s] += 1
				stackPairs[s] = 0
				sentences[len(sentences)-1].Pairs[s] += 1
				sentences[len(sentences)-1].Count++
			}
			sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			prev = unicode.ToLower(r)
		} else {
			prev = 0
			switch r {
			case '.':
				sentences = append(sentences, Sentence{
					Pairs:    map[string]float64{},
					Sentence: "",
					Count:    0,
				})
			case '\r', '\n':
			case '\t':
				sentences[len(sentences)-1].Sentence += fmt.Sprintf(" ")
			default:
				sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			}
		}
	}
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
			v, _ := s.Pairs[p]
			vs = append(vs, v)
		}
		log.Printf("%#v", vs)
		log.Printf("%#v", order)
		c.AddDataPair(fmt.Sprintf("Sentence %d: %s", i, s.Sentence), x, vs, co)
	}

	f, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const width = 1600
	height := 416
	i := image.NewRGBA(image.Rect(0, 0, width, height))
	igr := imgg.AddTo(i, 0, 0, width, height, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(i, i.Bounds(), bg, image.Point{}, draw.Src)
	c.Plot(igr)
	png.Encode(f, i)

}
