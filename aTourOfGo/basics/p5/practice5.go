package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("time.Now(): %v\n", t)
	t2 := time.Now().Weekday()
	fmt.Printf("time.Now().Weekday(): %v\n", t2)
}
