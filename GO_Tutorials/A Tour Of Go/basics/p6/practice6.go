package main

import (
	"fmt"
)

func main() {
	var p *int = new(int)
	*p = 5
	fmt.Printf("Value of p is %v\n", p)
	fmt.Printf("Value of *p is %v\n", *p)

	i := 67
	fmt.Printf("Value of &i is %v\n", &i)
}
