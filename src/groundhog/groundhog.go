package groundhog

import (
	"fmt"
	"utils"
	"os"
)

// Temperatures - values written in the input by the user
var Temperatures []float64

func getG() *float64 {
	if len(Temperatures) > Period {
		percentageIncrease := utils.GetIncreaseAverage(Temperatures, Period)
		res := &percentageIncrease
		return res
	}
	return nil
}

func getR() *int {
	if len(Temperatures) > Period {
		originalNumber := Temperatures[len(Temperatures) - Period - 1]
		newNumber := Temperatures[len(Temperatures) - 1]
		percentageIncrease := int(utils.GetEvolutionRate(originalNumber, newNumber))
		res := &percentageIncrease
		return res
	}
	return nil
}

func getS() *float64 {
	if len(Temperatures) >= Period {
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

	for ;; {
		fmt.Scanln(&input)

		if input == "STOP" {
			break
		}
		if !CheckTemperatureFormat(input) {
			os.Exit(84)
		}
		Temperatures = append(Temperatures, utils.ConvertStringToFloat(input))
		printTemperatureIncreaseAverage()
		printRelativeTemperatureEvolution()
		printStandardDeviation()
	}
}