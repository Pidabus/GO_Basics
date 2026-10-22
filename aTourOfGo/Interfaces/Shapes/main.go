package main

import (
	"math"
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{4, 5}
	circle := Circle{3}

	disp := calculateArea(rect)
	fmt.Println(disp)

	disp = calculateArea(circle)
	fmt.Println(disp)
}
