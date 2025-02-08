package utils

import "math"

func RoundToOneDecimalPlace(num float64) float64 {
	// Multiply by 10 to shift the decimal point
	shifted := num * 10.0
	// Round the shifted number
	rounded := math.Round(shifted)
	// Divide by 10 to shift the decimal point back
	return rounded / 10.0
}
