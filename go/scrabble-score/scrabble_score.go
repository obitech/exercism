package scrabble

import (
	"strings"
)

// Mapping maps letters to a point value.
type Mapping struct {
	Letters []string
	Points  int
}

// Scale is a slice of mappings to iterate over.
type Scale []Mapping

var s = Scale{
	Mapping{
		Letters: []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		Points:  1,
	},
	Mapping{
		Letters: []string{"D", "G"},
		Points:  2,
	},
	Mapping{
		Letters: []string{"B", "C", "M", "P"},
		Points:  3,
	},
	Mapping{
		Letters: []string{"F", "H", "V", "W", "Y"},
		Points:  4,
	},
	Mapping{
		Letters: []string{"K"},
		Points:  5,
	},
	Mapping{
		Letters: []string{"J", "X"},
		Points:  8,
	},
	Mapping{
		Letters: []string{"Q", "Z"},
		Points:  10,
	},
}

// BuildGrading constructs a map that can be used to query the amount of points
// for a given letter.
func (s Scale) BuildGrading() map[string]int {
	out := make(map[string]int)
	for _, mapping := range s {
		for _, letter := range mapping.Letters {
			out[letter] = mapping.Points
		}
	}
	return out
}

var grading = s.BuildGrading()

// Score scores a word in scrabble.
func Score(word string) int {
	var score int
	if word == "" {
		return score
	}

	for _, letter := range word {
		score += grading[strings.ToUpper(string(letter))]
	}
	return score
}
