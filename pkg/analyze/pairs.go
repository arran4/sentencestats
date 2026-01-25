package analyze

import (
	"fmt"
	"unicode"
)

type PairSentence struct {
	Pairs    map[string]float64
	Sentence string
	Count    int
}

func AnalyzePairs(s string) ([]PairSentence, map[string]float64) {
	allPairs := map[string]float64{}
	sentences := []PairSentence{{
		Pairs:    map[string]float64{},
		Sentence: "",
		Count:    0,
	}}
	var prev rune = 0
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			c := unicode.ToLower(r)
			if prev > 0 {
				p1, p2 := prev, c
				if p1 > p2 {
					p1, p2 = p2, p1
				}
				s := string([]rune{p1, p2})
				allPairs[s] += 1
				sentences[len(sentences)-1].Pairs[s] += 1
				sentences[len(sentences)-1].Count++
			}
			sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			prev = unicode.ToLower(r)
		} else {
			prev = 0
			switch r {
			case '.':
				sentences = append(sentences, PairSentence{
					Pairs:    map[string]float64{},
					Sentence: "",
					Count:    0,
				})
			case '\r', '\n':
				// Ignore newlines
			case '\t':
				sentences[len(sentences)-1].Sentence += " "
			default:
				sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			}
		}
	}

	// Remove the last sentence if it is empty, which happens if the input ends with a dot
	if len(sentences) > 0 && sentences[len(sentences)-1].Count == 0 && sentences[len(sentences)-1].Sentence == "" {
		sentences = sentences[:len(sentences)-1]
	}

	return sentences, allPairs
}
