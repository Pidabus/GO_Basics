package main

import (

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

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
}
