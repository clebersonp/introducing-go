package main

import "fmt"

func main() {

	// A map is an unordered collection of key-value pairs.
	// Maps are used to look up data quickly, given a key.

	// maps like slices or arrays, must be initialized before they can be used.
	// otherwise, they will result in a runtime error. (nil map)
	// a map of key type string and value type int
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println("Map. Length:", len(x), ", Value:", x["key"], ", Key-Value", x)

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y)

	// delete items from a map using the built-in delete function
	delete(x, "key")
	fmt.Println("Map. Length:", len(x), ", Value:", x["key"], ", Key-Value", x)

	// if a map has no key-value pairs, it will return the 'zero value' (which for string is the empty string, number 0, bool false)
	// We can use the comma-ok idiom to check if a key exists in a map
	value, ok := y[1]
	fmt.Println("ok for key number 1?", ok, ", Value:", value)

	// In Go, we often see code like this:
	if val, ok := y[1]; ok {
		fmt.Println("ok for key number 1?", ok, ", Value:", val)
	}
	// shorter way to create maps
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	// The outer map is used as a lookup table based on the element's symbol,
	// while the inner maps are used to store general information about the elements.
	// maps are often used like this.
	if el, ok := elements["Li"]; ok {
		fmt.Println(el, el["name"], el["state"])
	}

	v := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(v[2:5])

}
