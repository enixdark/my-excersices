package space

import "math"

type Planet string

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}

func Age(second float64, planet Planet) float64 {
    distance := Round(second / 31557600.0, 0.005)
	if planet == "Mercury" {
        return Round(distance / 0.2408467, 0.005)
	} else if planet == "Venus" {
        return Round(distance / 0.61519726, 0.005)
	} else if planet == "Mars" {
		return Round(distance / 1.8808158, 0.005)
	} else if planet == "Jupiter" {
		return Round(distance / 11.862615, 0.005)
	} else if planet == "Saturn" {
		return Round(distance / 29.447498, 0.005)
	} else if planet == "Uranus" {
		return Round(distance / 84.016846, 0.005)
	} else if planet == "Neptune" {
		return Round(distance / 164.79132, 0.005)
	} 
	return distance
}