package analyze

import (
	"testing"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedCount int
		expectedPairs map[string]float64
	}{
		{
			name:          "Simple pair",
			input:         "ab",
			expectedCount: 1,
			expectedPairs: map[string]float64{"ab": 1},
		},
		{
			name:          "Reverse pair",
			input:         "ba",
			expectedCount: 1,
			expectedPairs: map[string]float64{"ab": 1},
		},
		{
			name:          "Repeated pair",
			input:         "aba",
			expectedCount: 2,
			expectedPairs: map[string]float64{"ab": 2},
		},
		{
			name:          "Multiple pairs",
			input:         "abc",
			expectedCount: 2,
			expectedPairs: map[string]float64{"ab": 1, "bc": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentences, allPairs := Pairs(tt.input)

			if len(sentences) != 1 {
				t.Fatalf("expected 1 sentence, got %d", len(sentences))
			}

			s := sentences[0]
			if s.Count != tt.expectedCount {
				t.Errorf("sentence count = %d, want %d", s.Count, tt.expectedCount)
			}

			for pair, count := range tt.expectedPairs {
				if got := s.Pairs[pair]; got != count {
					t.Errorf("pair %s count = %f, want %f", pair, got, count)
				}
				if got := allPairs[pair]; got != count {
					t.Errorf("allPairs %s count = %f, want %f", pair, got, count)
				}
			}
		})
	}
}
