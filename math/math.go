package utils

import (
	"golang.org/x/exp/constraints"
	"math"
)

// PercentOf gets a certain percent out of a number.
// For example, 50% of 10 will return 5.
func PercentOf[Number constraints.Unsigned](a Number, percent Number) Number {
	return Number(math.Round(float64(a) * (float64(percent) / 100.0)))
}

// Min gets the smaller of the two values.
func Min[Number constraints.Unsigned](a Number, b Number) Number {
	if b > a {
		return a
	}
	return b
}

// Max gets the larger of the two values.
func Max[Number constraints.Unsigned](a Number, b Number) Number {
	if a > b {
		return a
	}
	return b
}
