package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	fmt.Printf("Type s: %[1]T, elements: %[1]v, lenght: %d, cap: %d\n", s, len(s), cap(s))
	for _, n := range s {
		n += 1
	}
	// slice s values has no changed. The range loop copies the values from the slice to a 'local variable' n.
	// Updating n will not affect the slice
	fmt.Println(s)

	// To update slice values:
	for i := range s {
		s[i] += 1
	}
	fmt.Println(s)

	// To update slice values:
	a, b, c := 1, 1, 1
	v := []*int{&a, &b, &c}
	fmt.Println(v)
	for _, n := range v {
		*n += 1
	}
	for _, n := range v {
		fmt.Printf("%v, %d\n", n, *n)
	}
}
