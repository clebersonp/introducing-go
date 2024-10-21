package main

import (
	"fmt"
	"os"
)

func fooWithType() error {
	var err *os.PathError = nil
	return err
}

// fooBestPractice to avoid that problem use a variable of type error instead.
// [named return value](/golang/named-return-values-parameters/)
// Best practice: Use the built-in error interface type,
// rather than a concrete type, to store and return error values.
func fooBestPractice() (err error) {
	return
}

func main() {
	err := fooWithType()
	// Why is nil not equal to nil in this example?
	// An interface value is equal to nil only if both its value and dynamic type are nil.
	// In the example below, fooWithType() return [nil, *os.PathError] and we compare it with [nil, nil].
	fmt.Printf(
		"fooWithType(). Type of error: %[1]T, value of error: %[1]v, Type of nil: %[2]T, value of nil: %[2]v, equals: %t\n",
		err, nil, err == nil,
	)

	err = fooBestPractice()
	fmt.Printf(
		"fooBestPractice(). Type of error: %[1]T, value of error: %[1]v, Type of nil: %[2]T, value of nil: %[2]v, equals: %t\n",
		err, nil, err == nil,
	)
}
