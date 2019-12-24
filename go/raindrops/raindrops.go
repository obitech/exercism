// Package raindrops implements a function to simulate raindrops.
package raindrops

import "strconv"

// Convert returns a different string, depending if input is evenly divisible
// through 3, 5 and/or 7.
func Convert(input int) string {
	by3, by5, by7 := input%3 == 0, input%5 == 0, input%7 == 0
	if by3 || by5 || by7 {
		var out string
		if by3 {
			out += "Pling"
		}
		if by5 {
			out += "Plang"
		}
		if by7 {
			out += "Plong"
		}
		return out
	}
	return strconv.Itoa(input)
}
