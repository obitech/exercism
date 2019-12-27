// Package triangle provides a function to determine the kind of a triangle.
package triangle

import (
	"math"
)

// Kind represents a kind of triangle.
type Kind int

const (
	// NaT is a Kind that doesn't represent a triangle.
	NaT Kind = iota
	// Equ is an equilateral Kind of triangle, having three equal sides.
	Equ
	// Iso is an isosceles Kind of triangle, having two equal sides.
	Iso
	// Sca is a scalene Kind of triangle, having inequal sides.
	Sca
)

// isTriangle checks if a given three sides, a valid triangle exists.
func isTriangle(a, b, c float64) bool {
	for _, v := range []float64{a, b, c} {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	if a+b+c <= 0 || a+b < c || a+c < b || b+c < a {
		return false
	}
	return true
}

// KindFromSides determines the kind of a triangle given the lenght of its sides.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if (a == b) || (a == c) || (b == c) {
		return Iso
	}
	return Sca
}
