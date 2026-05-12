package main

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	println("The area is: ", s.getArea())
}
