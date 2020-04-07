package groundhog

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	noArg = "There is no argument.\n"
	mustBeInteger = "The argument must be a positive integer.\n"
	mustBeGreatherThanTwo = "Impossible to detect a tendency if the argument is less than or equal to 2.\n"
	wrongTemperatureFormat = "Wrong temperature format, the format must be a decimal digit."
)

// Period - number of days (the argument)
var Period int = 0

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if (arg == "-h") {
			return true
		}
	}
	return false
}

// CheckTemperatureFormat - check period value given as argument with a regular expression
func CheckTemperatureFormat(input string) bool {
	var re = regexp.MustCompile("[-]?[0-9]{1,9}\\.[0-9]{1,1}")

	match := re.FindString(input)
	if len(match) != len(input) {
		fmt.Println(wrongTemperatureFormat)
		return false
	}
	return true;
}

func checkArg(arg string) bool {
	var re = regexp.MustCompile("[0-9]")

	match := re.FindAllString(arg, -1)
	if len(arg) != len(match) {
		fmt.Println(mustBeInteger)
		return false
	}
	resInt, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	if (resInt < 3) {
		fmt.Println(mustBeGreatherThanTwo)
		return false
	}
	Period = resInt
	return true
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println(noArg)
		return false
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println(tooManyArgs)
		return false
	}
	if (!checkArg(argsWithoutProg[0])) {
		return false
	}
	return true
}