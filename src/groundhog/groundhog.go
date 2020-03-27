package groundhog

import (
	"fmt"
	"utils"
)

// Temperatures - values written in the input by the user
var Temperatures []float64
var periodLoopIsFinished bool = false

func getG() *float64 {
	return nil
}

func getR() *int {
	if periodLoopIsFinished {
		originalNumber := utils.GetAverage(Temperatures)
		newNumber := Temperatures[len(Temperatures) - 1]
		percentageIncrease := int(utils.GetPercentageIncrease(originalNumber, newNumber))
		res := &percentageIncrease
		return res
	}
	return nil
}

func getS() *float64 {
	if periodLoopIsFinished {
		variance := utils.GetStandardDeviation(Temperatures, Period)
		res := &variance
		return res
	}
	return nil
}

func printTemperatureIncreaseAverage() {
	var g = getG()

	fmt.Printf("g=")
	if g == nil {
		fmt.Printf("nan")
	} else {
		fmt.Printf("%.2f", *g)
	}
}

func printRelativeTemperatureEvolution() {
	var r = getR()

	fmt.Printf("\t\tr=")
	if r == nil {
		fmt.Printf("nan%%")
	} else {
		fmt.Printf("%d%%", *r)
	}
}

func printStandardDeviation() {
	var s = getS()

	fmt.Printf("\t\ts=")
	if s == nil {
		fmt.Printf("nan")
	} else {
		fmt.Printf("%.2f", *s)
	}
	fmt.Printf("\n")
}

// Groundhog - main
func Groundhog() {
	var input string

	for i := 1; ; i++ {
		fmt.Scanln(&input)

		if input == "STOP" {
			break
		}
		if !CheckTemperatureFormat(input) {
			i--
			continue
		}
		if (i >= Period) {
			periodLoopIsFinished = true
		}
		Temperatures = append(Temperatures, utils.ConvertStringToFloat(input))
		printTemperatureIncreaseAverage()
		printRelativeTemperatureEvolution()
		printStandardDeviation()
	}
}