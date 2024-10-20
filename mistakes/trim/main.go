package main

import (
	"fmt"
	"strings"
)

func main() {
	// The Trim, TrimLeft and TrimRight functions strip all Unicode code points contained in a cutset.
	fmt.Println(strings.TrimRight("ABBA", "BA")) // Output: ""

	// To strip a trailing string, use strings.TrimSuffix function
	fmt.Println(strings.TrimSuffix("ABBA", "BA"))
}
