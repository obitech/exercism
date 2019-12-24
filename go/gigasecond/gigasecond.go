// Package gigasecond provides calculations around gigaseconds.
package gigasecond

import "time"

// AddGigasecond adds a Gigasecond and returns the final time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1E9 * time.Second)
}
