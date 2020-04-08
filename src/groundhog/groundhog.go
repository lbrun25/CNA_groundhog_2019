package groundhog

import (
	"fmt"
	"utils"
	"os"
)

const (
	impossibleToCalculate = "Impossible to calculate tendency switched and weirdest values without calculating all values over the given period."
	temperatureEqualToZero = "Impossible to calculate any value with a temperature strictly equal to zero."
	tab = "       "
)

// Temperatures - values written in the input by the user
var Temperatures []float64

// CountSwitchOccurs - count switches
var CountSwitchOccurs int

// MovingAverages - list of moving averages
var MovingAverages []float64

// StandardDeviations - list of standard deviations
var StandardDeviations []float64

// BollingerBands - list of Bollinger bands
var BollingerBands []float64

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
		StandardDeviations = append(StandardDeviations, *res)
		return res
	}
	return nil
}

func appendMovingAverage() {
	if len(Temperatures) >= Period {
		MovingAverages = append(MovingAverages, utils.GetAveragePeriod(Temperatures, Period))
	}
}

func appendBollingerBands() {
	if len(Temperatures) >= Period {
		BollingerBands = append(BollingerBands, utils.GetBollingerBand(Temperatures, MovingAverages, StandardDeviations))
	}
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

	fmt.Printf(tab)
	fmt.Printf("r=")
	if r == nil {
		fmt.Printf("nan%%")
	} else {
		fmt.Printf("%d%%", *r)
	}
}

func printStandardDeviation() {
	var s = getS()

	fmt.Printf(tab)
	fmt.Printf("s=")
	if s == nil {
		fmt.Printf("nan")
	} else {
		fmt.Printf("%.2f", *s)
	}
}

func printSwitchOccurs() {
	if len(Temperatures) > Period + 1 {
		if utils.IsSwitchOccurs(Temperatures, Period) {
			fmt.Printf(tab)
			fmt.Printf("a switch occurs")
			CountSwitchOccurs++
		}
	}
	fmt.Println("")
}

func printWeirdestValues() {
	weirdestValues := utils.GetWeirdestValues(Temperatures, BollingerBands, Period)

	for i := len(weirdestValues) - 1; i != -1; i-- {
		value := weirdestValues[i]
		fmt.Printf("%.1f", value)
		if i != 0 {
			fmt.Printf(", ")
		}
	}
	fmt.Println("]")
}

func printResults() {
	if (len(Temperatures) < Period) {
		fmt.Println(impossibleToCalculate)
		os.Exit(84)
	}
	fmt.Println("Global tendency switched", CountSwitchOccurs, "times")
	fmt.Printf("5 weirdest values are [")
	printWeirdestValues()
}

func appendTemperature(input string) {
	if !CheckTemperatureFormat(input) {
		os.Exit(84)
	}
	temperature := utils.ConvertStringToFloat(input)
	if temperature == 0 {
		fmt.Println(temperatureEqualToZero)
		os.Exit(84)
	} else {
		Temperatures = append(Temperatures, temperature)
	}
}

// Groundhog - main
func Groundhog() {
	var input string

	for ;; {
		fmt.Scanln(&input)

		if input == "STOP" {
			break
		}
		appendTemperature(input)
		printTemperatureIncreaseAverage()
		printRelativeTemperatureEvolution()
		printStandardDeviation()
		printSwitchOccurs()
		appendMovingAverage()
		appendBollingerBands()
	}
	printResults()
}