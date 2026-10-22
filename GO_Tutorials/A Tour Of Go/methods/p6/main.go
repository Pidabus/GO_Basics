package main

import (
	"fmt"
)

func main () {
	s := myString("Hello, World!")
	f := myFloat(3.9)

	fmt.Println(f, s)
}

type myString string 
type myFloat float64