package utils

import (
	"math"
)

// GetAveragePeriod - compute average on the period
func GetAveragePeriod(values []float64, period int) float64 {
	var sum float64 = 0.0

	for _, value := range values[len(values) - period:] {
		sum += value
	}
	return sum / float64(period)
}

// GetStandardDeviation - compute standard deviation on the period
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

// GetEvolutionRate - evolution rate on the period
func GetEvolutionRate(originalNumber float64, newNumber float64) int {
	increase := newNumber / originalNumber - 1
	res := math.Round(increase * 100)
	return int(res)
}

// GetIncreaseAverage - increase average on the period
func GetIncreaseAverage(values []float64, period int) float64 {
	var maxValues []float64

	currentPeriod := values[len(values) - period:]
	prevPeriod := values[len(values) - period - 1: len(values) - 1]
	for i, firstValue := range currentPeriod {
		if firstValue - prevPeriod[i] > 0 {
			maxValues = append(maxValues, firstValue - prevPeriod[i])
		} else {
			maxValues = append(maxValues, 0)
		}
	}
	return GetAveragePeriod(maxValues, period)
}