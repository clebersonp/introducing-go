package main

import (
	"fmt"
	"math"
)

// Shape
// Like a struct, an interface is created using the type keyword, followed by a name and the keyword interface.
// But instead of defining fields, we define a method set.
// Both Rectangle and Circle have area methods tha return float64, so both types implement the Shape interface.
// We can use interface types as arguments to functions.
// Nothing additional is required to implement an interface (there is no implements or extends keyword).
// It's sufficient to merely have a method with the same name and signature.
type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

// area implements the Shape interface because the signatures are the same
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// area implements the Shape interface because the signatures are the same
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	a := r.x1 + r.y1
	b := r.x2 + r.y2
	return 2 * (a + b)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// totalArea calculates the total of area for all type of Shapes(interface type)
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// MultiShape can have fields of interface type.
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{
		x: 0,
		y: 0,
		r: 5,
	}

	r := Rectangle{
		x1: 2,
		y1: 4,
		x2: 12,
		y2: 4,
	}

	fmt.Println("Total area:", totalArea(&c, &r))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{
				x: 0,
				y: 0,
				r: 5,
			},
			&Rectangle{
				x1: 4,
				y1: 10,
				x2: 20,
				y2: 4,
			},
		},
	}

	fmt.Println("MultiShape total area:", multiShape.area())
	fmt.Println("MultiShape total perimeter:", multiShape.perimeter())
}
