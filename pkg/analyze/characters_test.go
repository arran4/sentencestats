package analyze

import (
	"testing"
)

func TestAnalyzeCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []CharSentence
	}{
		{
			name:  "Single sentence",
			input: "abc",
			expected: []CharSentence{
				{
					Sentence: "abc",
					Count:    3,
					// Hist verification might be tedious, check specific indices
				},
			},
		},
		{
			name:  "Two sentences",
			input: "a.b.",
			expected: []CharSentence{
				{
					Sentence: "a",
					Count:    1,
				},
				{
					Sentence: "b",
					Count:    1,
				},
			},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: []CharSentence{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnalyzeCharacters(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("AnalyzeCharacters() got %d sentences, want %d", len(got), len(tt.expected))
				return
			}
			for i, s := range got {
				if s.Sentence != tt.expected[i].Sentence {
					t.Errorf("AnalyzeCharacters()[%d].Sentence = %v, want %v", i, s.Sentence, tt.expected[i].Sentence)
				}
				if s.Count != tt.expected[i].Count {
					t.Errorf("AnalyzeCharacters()[%d].Count = %v, want %v", i, s.Count, tt.expected[i].Count)
				}
			}
		})
	}
}

func TestAnalyzeCharacters_Hist(t *testing.T) {
	s := "AaBbC"
	got := AnalyzeCharacters(s)
	if len(got) != 1 {
		t.Fatalf("expected 1 sentence, got %d", len(got))
	}

	// 'a' count (index 0) should be 2 (one 'A', one 'a')
	if got[0].Hist[0] != 2 {
		t.Errorf("expected 'a' count to be 2, got %f", got[0].Hist[0])
	}
	// 'b' count (index 1) should be 2
	if got[0].Hist[1] != 2 {
		t.Errorf("expected 'b' count to be 2, got %f", got[0].Hist[1])
	}
	// 'c' count (index 2) should be 1
	if got[0].Hist[2] != 1 {
		t.Errorf("expected 'c' count to be 1, got %f", got[0].Hist[2])
	}
}
