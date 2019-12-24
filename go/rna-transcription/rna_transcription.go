// Package strand defines a DNA to RNA conversion function.
package strand

import "strings"

// complement defines nucleotide complements.
var complement = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA converts a DNA string to its RNA complement.
func ToRNA(dna string) string {
	var rna string
	for _, r := range dna {
		rna += complement[strings.ToUpper(string(r))]
	}
	return rna
}
