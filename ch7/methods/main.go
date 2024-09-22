package main

import (
	"fmt"
	"math"
)

// Circle is a struct of type Circle
type Circle struct {
	x, y, r float64
}

// Methods are special type of function that we can perform with a specific structure.
// area performs the circle area
// between the keyword func and the name of the function, we've added a 'receiver'.
// receiver is like a parameter - it has a name and a type - but by creating
// the function in this way, it allows us to call the function using the '.' operator to circle structure.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func main() {

	c := Circle{
		x: 0,
		y: 0,
		r: 2,
	}

	// call the method with '.' operator
	fmt.Println("Circle c:", c, ", area:", c.area())

}
