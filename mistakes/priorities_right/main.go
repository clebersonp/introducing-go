package main

import (
	"fmt"
)

func main() {
	n := 43210 // time in seconds
	fmt.Println(n/60*60, "hours and", n%60*60, "seconds")

	// The *, /, and % have the same precedence and are evaluated left to right:
	// n/60*60 is the same as (n/60)*60
	// Insert a pair of parentheses to force the correct evaluation order.
	fmt.Println(n/(60*60), "hours and", n%(60*60), "seconds")

	// Or better yet, use a constant
	const secPerHour = 60 * 60
	fmt.Println(n/secPerHour, "hours and", n%secPerHour, "seconds")
}
