package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("time.Now(): %v\n", t)
	fmt.Printf("time.Now().Hour(): %v\n", t.Hour())
	fmt.Printf("time.Now().Weekday(): %v\n", t.Weekday())
}
