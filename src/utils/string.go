package utils

import (
	"strconv"
)

// ConvertStringToFloat - convert a string to float by checking error
func ConvertStringToFloat(str string) float64 {
	s, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(err)
	}
	return s
}