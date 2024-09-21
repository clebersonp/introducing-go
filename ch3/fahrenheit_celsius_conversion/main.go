package main

import "fmt"

func main() {
	fmt.Print("Enter Fahrenheit to conversion: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("Fahrenheit:", fahrenheit, "Celsius:", celsius)
}
