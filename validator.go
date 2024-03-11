package bmicalculator

import (
	"fmt"
	"strconv"
)

func ParseFloat(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, fmt.Errorf("failed converting to float: %v", err)
	}

	return num, nil
}

func IsValidHeight(meters float64) bool {
	return meters > 0 && meters < 3.1
}

func IsValidWeight(kilograms float64) bool {
	return kilograms > 0 && kilograms < 386
}
