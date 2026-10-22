package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) String() string { // Doesn't mutate values, so a pointer receiver is not required.
	return fmt.Sprintf("Counter: %d", c.value)
}

func main() {
	c := &Counter{}
	fmt.Println(c)
	c.Increment()
	fmt.Println(c) // Prints "Counter: 1"
	c.Increment()
	fmt.Println(c)
}
