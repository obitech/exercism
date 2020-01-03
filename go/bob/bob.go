// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

func forceful(s string) bool {
	s = s[1:]
	for _, c := range s {
		if unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// String is empty.
	if len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	}

	// Remove all special characters, numbers, etc.
	re := regexp.MustCompile(`[\d,\s\%\-_:).!'^*@#$(*]`)
	remark = re.ReplaceAllString(remark, "")

	// String contains characters and ends in ?
	if len(remark) > 0 && remark[len(remark)-1] == '?' {
		// String contains a valid letter and might be forceful
		if len(remark) > 1 && forceful(remark) {
			return "Calm down, I know what I'm doing!"
		}
		// String contains no letter
		return "Sure."
	}

	// String contains letters and might be forceful
	if len(remark) > 0 && forceful(remark) {
		return "Whoa, chill out!"
	}

	// Everything else
	return "Whatever."
}
