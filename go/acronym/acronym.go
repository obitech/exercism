// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate returns the acronym for a given string
func Abbreviate(s string) string {
	if s == "" {
		return s
	}
	var out string
	t := strings.Split(s, " ")
	for _, c := range t {
		if c != "" {
			c = strings.Trim(c, "_")
			if strings.Contains(c, "-") {
				sub := strings.Split(c, "-")
				for _, sc := range sub {
					if len(sc) == 0 {
						continue
					}
					out += strings.ToUpper(string(sc[0]))
				}
				continue
			}
			out += strings.ToUpper(string(c[0]))
		}
	}
	return out
}
