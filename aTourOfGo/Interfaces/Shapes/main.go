package main

type Shapes interface {
	Area() float64
}

type Rectangle struct {
	area, width float64
}

type Circle struct {
	radius float64
}

func main() {
}
