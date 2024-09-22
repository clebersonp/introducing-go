package main

import "fmt"

// half halves and returns true if it was even or false if it was odd.
func half(i int) (int, bool) {
	switch i % 2 {
	case 0:
		return 1, true
	default:
		return 0, false
	}
}

// greatest finds the greatest number in a list of numbers (variadic parameter).
func greatest(args ...int) int {
	big := 0
	for _, n := range args {
		if n > big {
			big = n
		}
	}
	return big
}

// makeOddGenerator returns a function to generate odd numbers using closure to persist the next value and return the current.
func makeOddGenerator() func() int {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	value, status := half(1)
	fmt.Println("half(1):", value, status)

	value, status = half(2)
	fmt.Println("half(2):", value, status)

	fmt.Println()

	args := []int{2, 8, 99, 102, 45, 67, 1}
	fmt.Println("List of numbers:", args, ", greatest number:", greatest(args...))

	fmt.Println()

	nextOdd := makeOddGenerator()
	fmt.Println("Next Odd:", nextOdd()) // 1
	fmt.Println("Next Odd:", nextOdd()) // 3
	fmt.Println("Next Odd:", nextOdd()) // 5

	fmt.Println()

	fmt.Println("Fibonacci of 15:", fib(15))

	fmt.Println()

	x := 1.5
	square(&x)
	fmt.Println("Square of 1.5:", x)

	a, b := 5, 10
	fmt.Println("Before swapping. a:", a, "b:", b)
	swap(&a, &b)
	fmt.Println("After swapping. a:", a, "b:", b)

}
