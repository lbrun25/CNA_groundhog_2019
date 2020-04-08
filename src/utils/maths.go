package utils

import (
	"math"
	"sort"
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

// IsSwitchOccurs - is a switch occurs on the period ?
func IsSwitchOccurs(values []float64, period int) bool {
	first := values[len(values) - 1]
	second := values[len(values) - period - 1]
	currentEstimation := math.Round((first / second - 1) * 100)

	first = values[len(values) - 2]
	second = values[len(values) - period - 2]
	prevEstimation := math.Round((first / second - 1) * 100)

	sumAbs := math.Abs(prevEstimation + currentEstimation) 
	if sumAbs != math.Abs(prevEstimation) + math.Abs(currentEstimation) {
		return true
	}
	return false
}

// GetBollingerBand - Bollinger band
func GetBollingerBand(values []float64, movingAverages []float64, standardDeviations []float64) float64 {
	lastMovingAverage := movingAverages[len(movingAverages) - 1]
	lastStandardDeviation := standardDeviations[len(standardDeviations) - 1]
	lastValue := values[len(values) - 1]

	upperBand := lastMovingAverage + 2 * lastStandardDeviation
	lowerBand := lastMovingAverage - 2 * lastStandardDeviation
	res := (lastValue - lowerBand) / (upperBand - lowerBand)
	return res
}

// GetWeirdestValues - get weirdest values on the period
func GetWeirdestValues(values []float64, bollingerBands []float64, period int, limit int) []float64 {
	var absList []float64
	var sortedAbsList []float64
	var resList []float64

	for _, value := range bollingerBands {
		absList = append(absList, math.Abs(value - float64(limit) / 10))
	}
	sortedAbsList = make([]float64, len(absList))
	copy(sortedAbsList, absList)
	sort.Float64s(sortedAbsList)
	sortedAbsList = sortedAbsList[len(sortedAbsList) - limit:]
	for _, value := range sortedAbsList {
		index := SliceFloatIndex(absList, value) + period - 1
		resList = append(resList, values[index])
	}
	return resList
}