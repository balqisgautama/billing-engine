package utils

import "time"

func CalculateDifferenceWeeks(start, end time.Time) int {
	// Calculate the difference in days
	difference := end.Sub(start)

	// Convert the difference to weeks
	weeks := int(difference.Hours() / (24 * 7))
	return weeks
}

func CalculateDifferenceWeeksFloat(start, end time.Time) float64 {
	// Calculate the difference in days
	difference := end.Sub(start)

	// Convert the difference to weeks
	weeks := difference.Hours() / (24 * 7)
	// only return 1 decimal places
	return RoundToOneDecimalPlace(weeks)
}
