package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// This means that type T implements interface I
func (t T) M() {
	fmt.Println(t.S)
}

func main () {
	// var i I = T{"Hello"}
	t := T{"Hello I'm I"}
	i := T{"Hello"}

	t.M()
	i.M()
}