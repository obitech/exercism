package etl

import "strings"

// Transform transform an old Scrabble point system to a newer one.
func Transform(old map[int][]string) map[string]int {
	new := make(map[string]int)
	for point, letters := range old {
		for _, l := range letters {
			new[strings.ToLower(l)] = point
		}
	}
	return new
}
