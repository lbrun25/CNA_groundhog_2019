package groundhog

import (
	"fmt"
)

func getG() *float32 {
	return nil
}

func getR() *int {
	return nil
}

func getS() *float32 {
	return nil
}

func printTemperatureIncreaseAverage() {
	var g = getG()

	fmt.Printf("g=")
	if g == nil {
		fmt.Printf("nan")
	} else {
		fmt.Printf("%d", g)
	}
}

func printRelativeTemperatureEvolution() {
	var r = getR()

	fmt.Printf("\t\tr=")
	if r == nil {
		fmt.Printf("nan")
	} else {
		fmt.Printf("%d", r)
	}
}

func printStandardDeviation() {
	var s = getS()

	fmt.Printf("\t\ts=")
	if s == nil {
		fmt.Printf("nan")
	} else {
		fmt.Printf("%d", s)
	}
	fmt.Printf("\n")
}

func periodLoop() {
	for i := 0; i < Period; i++ {
		var input string
		fmt.Scanln(&input)
		if !CheckTemperatureFormat(input) {
			i--;
			continue;
		}
		printTemperatureIncreaseAverage()
		printRelativeTemperatureEvolution()
		printStandardDeviation()
	}
}

// Groundhog - main
func Groundhog() {
	periodLoop();
	// To do: Detect a tendency and compute values
}