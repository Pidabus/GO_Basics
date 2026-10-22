package main

import (
	"math"
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

}