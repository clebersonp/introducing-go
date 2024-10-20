package main

import (
	"encoding/json"
	"fmt"
)

type PersonWithoutExportFields struct {
	name string
	age  int
}

type PersonWithExportedFields struct {
	Name string
	Age  int
}

type PersonWithJsonTag struct {
	Name string `json:"my_name"`
	Age  int    `json:"my_age"`
}

func main() {
	p := PersonWithoutExportFields{
		name: "Alice",
		age:  22,
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	// Why json.Marshal produces empty struct data?
	fmt.Printf("Type of p: %T, state of p: %v\n", p, string(jsonData))

	// json.Marshal only produces the output to 'exported' fields of a Go struct

	p2 := PersonWithExportedFields{
		Name: "Alice",
		Age:  22,
	}
	jsonData, err = json.Marshal(p2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of p2: %T, state of p2: %v\n", p2, string(jsonData))

	// Using json tags in struct fields we can change the serialized struct field's name without change the struct
	p3 := PersonWithJsonTag{
		Name: "Alice",
		Age:  22,
	}
	jsonData, err = json.Marshal(p3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of p3: %T, state of p3: %v\n", p3, string(jsonData))

	// json.Marshal with map
	person := map[string]any{
		"first_name": "Alice",
		"age":        22,
	}
	jsonData, err = json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of person: %T, state of person: %v\n", person, string(jsonData))
}
