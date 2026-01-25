package analyze

import (
	"fmt"
	"unicode"
)

type CharSentence struct {
	Hist     [26]float64
	Sentence string
	Count    int
}

func AnalyzeCharacters(s string) []CharSentence {
	sentences := []CharSentence{{}}
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			c := unicode.ToLower(r)
			sentences[len(sentences)-1].Hist[c-'a'] += 1
			sentences[len(sentences)-1].Sentence += fmt.Sprintf("%c", r)
			sentences[len(sentences)-1].Count++
		} else {
			switch r {
			case '.':
				sentences = append(sentences, CharSentence{})
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

	return sentences
}
