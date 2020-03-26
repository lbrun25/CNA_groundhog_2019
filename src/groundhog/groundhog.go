package groundhog

import (
	"fmt"
	"os"
	"strconv"
	"utils"
)

// Temperatures - values written in the input by the user
var Temperatures []float64
var periodLoopIsFinished bool = false

func getG() *float64 {
	return nil
}

func getR() *int {
	return nil
}

func getS() *float64 {
	if periodLoopIsFinished {
		variance := utils.GetStandardDeviation(Temperatures)
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

func handleStop(input string) {
	if input == "STOP" {
		os.Exit(0)
	}
}

func convertStringToFloat(str string) float64 {
	s, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(err)
	}
	return s
}

// Groundhog - main
func Groundhog() {
	var input string

	for i := 0; ; i++ {
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
		Temperatures = append(Temperatures, convertStringToFloat(input))
		printTemperatureIncreaseAverage()
		printRelativeTemperatureEvolution()
		printStandardDeviation()
	}
}