// Package space provides the implementation of 'Age' exercise.
package space

import (
	"math"
)

// Planet type represents the planet's name
type Planet string

// Duration in seconds of a year on Earth
const earthYear = 31557600

// Age returns how old someone would be on a given planet
func Age(ageInSeconds float64, planet Planet) float64 {
	var age float64

	switch planet {
	case "Earth":
		age = ageInSeconds / earthYear
	case "Mercury":
		age = ageInSeconds / (earthYear * 0.2408467)
	case "Venus":
		age = ageInSeconds / (earthYear * 0.61519726)
	case "Mars":
		age = ageInSeconds / (earthYear * 1.8808158)
	case "Jupiter":
		age = ageInSeconds / (earthYear * 11.862615)
	case "Saturn":
		age = ageInSeconds / (earthYear * 29.447498)
	case "Uranus":
		age = ageInSeconds / (earthYear * 84.016846)
	case "Neptune":
		age = ageInSeconds / (earthYear * 164.79132)
	}

	return math.Round(age*100) / 100
}
