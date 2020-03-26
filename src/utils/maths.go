package utils

import (
	"math"
)

// GetAverage - compute average
func GetAverage(values []float64) float64 {
	var res float64 = 0.0

	for _, value := range values {
		res += value
	}
	return res / float64(len(values))
}

// GetStandardDeviation - compute variance
func GetStandardDeviation(values []float64) float64 {
	var res float64 = 0.0
	var squareValues float64 = 0.0
	average := GetAverage(values)

	for _, value := range values {
		squareValues += math.Pow(value - average, 2)
	}
	res = math.Sqrt(squareValues / 7)
	return res
}

// GetPercentageIncrease - compute percentage increase
func GetPercentageIncrease(originalNumber float64, newNumber float64) float64 {
	increase := newNumber - originalNumber
	res := increase / originalNumber * 100
	return res
}

// GetPercentageDecrease - compute percentage decrease
func GetPercentageDecrease(originalNumber float64, newNumber float64) float64 {
	decrease := originalNumber - newNumber
	res := decrease / originalNumber * 100
	return res
}