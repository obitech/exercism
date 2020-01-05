// Package isogram provides a function to check if a
// word is an isogram: a word or phrase without a repeating letter.
package isogram

import "unicode"

func skip(r rune) bool {
	if r == ' ' || r == '-' {
		return true
	}
	return false
}

// IsIsogram checks if a string is an isogram.
func IsIsogram(s string) bool {
	if s == "" {
		return true
	}
	check := make(map[rune]bool)
	for _, r := range s {
		r = unicode.ToLower(r)
		if !skip(r) {
			if _, ok := check[r]; ok {
				return false
			}
			check[r] = true
		}
	}
	return true
}
