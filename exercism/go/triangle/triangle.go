// Package triangle space provides the implementation of 'Triangle' exercise.
package triangle

import (
	"math"
)

// Kind represents the kind of the triangle.
type Kind int

const (
	//NaT : Not a tringle
	NaT = iota
	//Equ : Equilateral
	Equ
	//Iso : Isosceles
	Iso
	//Sca : Scalene
	Sca
)

// KindFromSides returns the kind of a triangle given the size of it's sides.
func KindFromSides(a, b, c float64) Kind {
	if !isTiangle(a, b, c) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

func isTiangle(a, b, c float64) bool {
	if !isValidSide(a) || !isValidSide(b) || !isValidSide(c) {
		return false
	}

	if a+b < c || a+c < b || b+c < a {
		return false
	}
	return true
}

func isValidSide(x float64) bool {
	if x <= 0 || math.IsNaN(x) || math.IsInf(x, 1) || math.IsInf(x, -1) {
		return false
	}
	return true
}