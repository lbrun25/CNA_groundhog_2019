package groundhog

import (
	"fmt"
	"utils"
	"os"
)

const (
	impossibleToCalculate = "Impossible to calculate tendency switched and weirdest values without calculating all values over the given period."
	tab = "       "
)

// Temperatures - values written in the input by the user
var Temperatures []float64
// CountSwitchOccurs - count switches
var CountSwitchOccurs int

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

func printResults() {
	if (len(Temperatures) < Period) {
		fmt.Println(impossibleToCalculate)
		os.Exit(84)
	}
	fmt.Println("Global tendency switched", CountSwitchOccurs, "times")
	printWeirdestValues()
}

func printWeirdestValues() {

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
		printSwitchOccurs()
	}
	printResults()
}