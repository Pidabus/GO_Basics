package main

import (
	"fmt"
	"time"
)

func sum(s []int) int{
	time.Sleep(5000 * time.Millisecond)
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	x := sum(s[:len(s)/2])
	y := sum(s[len(s)/2:])

	fmt.Println(x, y)
}