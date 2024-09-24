package mathy

import "testing"

// Go includes a special program that makes writing tests easier: 'go test'
// The Go compiler knows to ignore code in any files that end with _test.go,
// so the code defined in this file is only used by 'go test' (and not 'go install' or 'go build').

// Special package for testing called 'testing'
// We must define functions to test starting with the word 'Test' followed by the name of the function we want to test.
// Test functions takes a special argument of *testing.T type.

// To run the tests: go test ./ch9/mathy/
// or verbose with: go test -v ./ch9/mathy/

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got", v)
	}
}

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

// To run a specific test:
// go test -v ./ch9/mathy/ -run TestAverageWithTableDrivenTests

func TestAverageWithTableDrivenTests(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Errorf("Average(%v) = %f; want %f", pair.values, v, pair.average)
		}
	}
}
