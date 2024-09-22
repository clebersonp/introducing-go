package main

import (
	"fmt"
	"math"
	"reflect"
)

// Circle is a struct of type Circle
type Circle struct {
	x, y, r float64
}

// function accessing Circle fields
// when a function needs to modify any fields (state) of a struct
// we need create a function with pointer parameter instead.
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// Initialization and set fields to default zero value.
	// 0 for int, 0.0 for float, "" for string, nil for pointer, false for bool, etc.
	var c Circle
	fmt.Println("Type of c:", reflect.TypeOf(c), "value:", c)

	// Initialization using the new function (pointer). (new in this way is somewhat uncommon).
	c1 := new(Circle)
	fmt.Println("Type of c1:", reflect.TypeOf(c1), "value:", c1)

	// Give each of the fields an initial value.
	c2 := Circle{
		x: 0,
		y: 0,
		r: 0,
	}
	fmt.Println("Type of c2:", reflect.TypeOf(c2), "value:", c2)

	// leave off the field names if we know the order they were defined
	c3 := Circle{0, 0, 5}
	fmt.Println("Type of c3:", reflect.TypeOf(c3), "value:", c3)

	// or using pointer to the struct
	c4 := &Circle{0, 0, 9}
	fmt.Println("Type of c4:", reflect.TypeOf(c4), "value:", c4)

	// we can access fields using the '.' operator:
	fmt.Println("c4.x:", c4.x, "c4.y:", c4.y, "c4.r:", c4.r)

	fmt.Println("Circle c3:", c3, "Area:", circleArea(c3))
}
