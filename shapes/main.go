package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{base: 5, height: 3}
	fmt.Println("\nPRESENT TRIANGLE VALUES")
	showShape(tr)

	fmt.Println("\n PRESENT SQUARE")
	sq := square{4}
	showShape(sq)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func showShape(s shape) {
	fmt.Println(s.getArea())
}
