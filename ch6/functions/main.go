package main

import "fmt"

func main() {

	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println("Average of", xs, "is", average(xs))
	i1, i2 := f()
	fmt.Println("Call to f():", i1, i2)
	fmt.Println("Call to f2():", f2())
}

// average computes the average of a series of numbers.
// functions start with the keyword func, followed by the function's name. The parameters (input) of the function
// are defined like this: name type, name type, ... After the parameters, we put the return type. Collectively, the
// parameters and the return type are known as the function's signature.
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// f multiple values can be returned
// Multiple values are often used to return an error value along with the result
// (x, err := f()), or a boolean to indicate success (x, ok := f())
func f() (int, int) {
	return 5, 6
}

// f2 return types can have names
func f2() (r int) {
	r = 1
	return
}
