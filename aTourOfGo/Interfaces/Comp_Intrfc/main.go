package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func describeShape(g Geometry) {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}

func main() {
	rect := Rectangle{5, 7}
	describeShape(rect)
}
