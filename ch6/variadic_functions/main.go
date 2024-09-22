package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3))

	xs := []int{1, 2, 3}
	// we can pass a slice as a variadic parameter using the syntax: <slice_name>...
	// That form is similar to fmt.Println(a ...any) function that receives a variadic parameter of any type
	fmt.Println(add(xs...))
}

// add function use a special form available for the last parameter in Go called 'variadic parameter'.
// Variadic parameter must be the last parameter in a function list parameter.
// syntax: <variable_name> ...<type>
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
