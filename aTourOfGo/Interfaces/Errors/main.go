package main

import (
	"math"
	"fmt"
)

type CalculationError struct {
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{
			msg: "Invalid Input",
		}
	}
	return math.Sqrt(val), nil
}

func main() {
	x, err := performCalculation(5)
	fmt.Println(x, err)

	x, err = performCalculation(-3)
	fmt.Println(x, err)
}