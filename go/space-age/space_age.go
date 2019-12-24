// Package space provides a for calculating a person's age on a planet.
package space

import (
	"strings"
)

// Planet is a type indicating a planet in the solar system.
type Planet string

// Age takes the age (in seconds) of a person living on earth and calculates
// how old (in years) that person would be on any planet in the solar system.
func Age(age float64, planet Planet) float64 {
	// Age on earth in years.
	onEarth := age / 31557600
	switch strings.ToLower(string(planet)) {
	case "earth":
		return onEarth
	case "mercury":
		return onEarth / 0.2408467
	case "venus":
		return onEarth / 0.61519726
	case "mars":
		return onEarth / 1.8808158
	case "jupiter":
		return onEarth / 11.862615
	case "saturn":
		return onEarth / 29.447498
	case "uranus":
		return onEarth / 84.016846
	case "neptune":
		return onEarth / 164.79132
	}
	return 0
}
