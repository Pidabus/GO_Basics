package main

import "fmt"

type Greeter interface {
	Greet() string
}

type English struct{}

func (English) Greet() string {
	return "hello"
}

func main() {
	var g Greeter = English{}
	fmt.Println(g.Greet())
}
