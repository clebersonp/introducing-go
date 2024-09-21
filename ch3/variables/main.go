package main

import (
	"fmt"
	"reflect"
)

func main() {
	// variable can be declared with var keyword
	// assigning literal string value to variable x
	// declaration: <var_keyword> <variable_name> <data_type> = <value>
	var x string = "Hello World"

	// using variable to get the value instead of using string literals
	fmt.Println(x)

	// using two statements to declare and assign value
	var y string
	// assignment: <variable_name> = <value>
	y = "Hello World Again"
	fmt.Println(y)

	// concatenating strings
	x = x + " " + y + "!"
	fmt.Println(x)

	// special assignment operator
	x += "!!"
	fmt.Println(x)

	var h string = "Hello"
	var w string = "World"
	// comparing strings and return a boolean value
	fmt.Println(h+" == "+w+"?", h == w)

	var h2 string = "Hello"
	fmt.Println(h+" == "+h2+"?", h == h2)

	// creating a new variable with a starting value with shorter statement using operator :=
	// type is not required as it is inferred from the literal value by compiler
	z := "Hello World"
	fmt.Println(z)

	// declare and initialize a variable in one line without explicit type
	var t = "Hello World"
	fmt.Println(t)

	// declare and initialize a variable in one line without explicit type and type
	// generally, you should use this shorter form whenever possible
	v := 5
	fmt.Println(v, "of type", reflect.TypeOf(v))
}
