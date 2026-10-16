package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Record the exact starting time
	start := time.Now()

	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// 2. Calculate the difference between "then" and "now"
	duration := time.Since(start)

	// 3. Print the elapsed time
	fmt.Printf("\nExecution time: %v\n", duration)
}