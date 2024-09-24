package mathy

// In Go, if something starts with a capital letter, that means other package (and programs) are able to see it and
// use it. We call 'export a function'
// If the function name was lower letter like 'average', only functions over the same package or same file can use it.
// Package names match the folders they fall in. It's a lot easier if you stay within this pattern.

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
