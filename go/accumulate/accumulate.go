// Package accumulate performs repeated conversions on a slice of string.
package accumulate

// Accumulate applies a converter function to each string in a slice of
// strings.
func Accumulate(input []string, converter func(string) string) []string {
	if len(input) == 0 {
		return input
	}

	var out = make([]string, len(input))
	for idx, s := range input {
		out[idx] = converter(s)
	}
	return out
}
