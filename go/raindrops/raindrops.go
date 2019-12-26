// Package raindrops implements a function to simulate raindrops.
package raindrops

import "strconv"

// Convert returns a different string, depending if input is evenly divisible
// through 3, 5 and/or 7.
func Convert(input int) string {
	var out string
	if input%3 == 0 {
		out += "Pling"
	}
	if input%5 == 0 {
		out += "Plang"
	}
	if input%7 == 0 {
		out += "Plong"
	}
	if out == "" {
		return strconv.Itoa(input)
	}
	return out
}
