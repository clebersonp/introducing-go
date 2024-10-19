package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	// the uninitialized pointer p in the main function is nil
	// pointer must be initialized before use it.
	//var p *Point // panic: runtime error: invalid memory address or nil pointer dereference
	p := new(Point)
	p.X = 3
	p.Y = 4
	fmt.Println(p.Abs())

	// Since methods with pointer receivers take either a value or a pointer, you could also skip the pointer altogether
	var point Point // zero value: Point{0, 0}
	point.X = 2
	point.Y = 5
	fmt.Println(point.Abs())
}
