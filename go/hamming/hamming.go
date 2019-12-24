// Package hamming provides functionality to calculate the hamming distance
// between strings.
package hamming

import "fmt"

// Distance calculates the hamming distance of two equal length strings.
func Distance(a, b string) (int, error) {
	la := len(a)
	lb := len(b)
	if la != lb {
		return 0, fmt.Errorf("arguments have different lengths: %d != %d", la, lb)
	}

	var d int
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
