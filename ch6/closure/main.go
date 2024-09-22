package main

import (
	"fmt"
	"reflect"
)

func main() {

	// we can create a function inside of functions.
	// add is a local variable of type func(int, int) int
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("Type of 'add' variable:", reflect.TypeOf(add), add(1, 1))

	// local function also has access to other local variables
	// a function like this together with the nonlocal variables it references is known as a closure.
	// In this case, 'increment' and the variable 'x' form the closure.
	x := 0
	increment := func() int {
		x++
		return x
	}

	// to invoke a local function we must call with parentheses and its list of parameters
	increment()
	increment()
	fmt.Println("Type of 'increment':", reflect.TypeOf(increment), ", Value:", x)

	fmt.Println()

	// One way to use closure is by writing a function that returns another function,
	// which when called, can generate a sequence of numbers.
	// I'm not invoking the inside function, only return the type function
	nextEven := makeEvenGenerator()
	// Here I'm calling the inside function to be evaluated
	fmt.Println("Type of 'nextEven':", reflect.TypeOf(nextEven), ", Value:", nextEven()) // 0
	fmt.Println("Type of 'nextEven':", reflect.TypeOf(nextEven), ", Value:", nextEven()) // 2
	fmt.Println("Type of 'nextEven':", reflect.TypeOf(nextEven), ", Value:", nextEven()) // 4

	otherEvent := makeEvenGenerator()
	fmt.Println("Type of 'otherEvent':", reflect.TypeOf(otherEvent), ", Value:", otherEvent()) // 0
	fmt.Println("Type of 'otherEvent':", reflect.TypeOf(otherEvent), ", Value:", otherEvent()) // 2
	fmt.Println("Type of 'otherEvent':", reflect.TypeOf(otherEvent), ", Value:", otherEvent()) // 4

	// Each time makeEvenGenerator was called, will be returned a new reference of that function.
	// So, each reference will be its state
}

// 'makeEvenGenerator' returns a function that generates even number.
// Each time it's called (inside function), it adds 2 to the local i variable,
// which - unlike normal local variables - persists between calls.
func makeEvenGenerator() func() uint {
	// variable 'i' and function that access it form the closure.
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
