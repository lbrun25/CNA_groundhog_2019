package utils

import (
	"math"
)

// GetAveragePeriod - compute average
func GetAveragePeriod(values []float64, period int) float64 {
	var sum float64 = 0.0

	for _, value := range values[len(values) - period:] {
		sum += value
	}
	return sum / float64(period)
}

// GetStandardDeviation - compute variance
func GetStandardDeviation(values []float64, period int) float64 {
	var res float64 = 0.0
	var squareValues float64 = 0.0
	average := GetAveragePeriod(values, period)

	for _, value := range values[len(values) - period:] {
		squareValues += math.Pow(value - average, 2)
	}
	res = math.Sqrt(squareValues / float64(period))
	return res
}

// GetPercentageIncrease - compute percentage increase
func GetPercentageIncrease(originalNumber float64, newNumber float64) int {
	increase := newNumber / originalNumber - 1
	res := math.Round(increase * 100)
	return int(res)
}

// GetPercentageDecrease - compute percentage decrease
func GetPercentageDecrease(originalNumber float64, newNumber float64) float64 {
	decrease := originalNumber - newNumber
	res := decrease / originalNumber * 100
	return res
}