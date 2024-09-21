package main

import "fmt"

func main() {
	printAllNumbersDivisibleBy3()
	fmt.Println()
	printAllNumbersMultiplesOf3OrFive()
}

// printAllNumbersDivisibleBy3 prints out all the numbers between 1 and 100 that are
// evenly divisible by 3 (i.e., 3, 6, 9, etc.)
func printAllNumbersDivisibleBy3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// printAllNumbersMultiplesOf3OrFive prints the numbers from 1 to 100, but for multiples of
// three, print “Fizz” instead of the number, and for the multiples of five, print
// “Buzz.” For numbers that are multiples of both three and five, print “FizzBuzz.”
func printAllNumbersMultiplesOf3OrFive() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
