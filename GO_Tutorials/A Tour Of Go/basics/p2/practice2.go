package main

import (
	"fmt"
	"math"
)
	
func sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprint(math.Sqrt(-x)) + "i"
	}
	return "33"
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
