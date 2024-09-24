package main

import (
	"fmt"
	"sort"
)

// The sort package contains functions for sorting arbitrary data

type Person struct {
	Name string
	Age  int
}

// ByName implements sort.Interface
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// ByAge implements sort.Interface
type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{
			Name: "Jack",
			Age:  10,
		},
		{
			Name: "Dan",
			Age:  3,
		},
		{
			Name: "Jill",
			Age:  9,
		},
	}
	// sort by name
	// casting kids to underlying type of []Person (ByName)
	sort.Sort(ByName(kids))
	fmt.Println("Sort by name:", kids)

	// sort by age
	// casting kids to underlying type of []Person (ByAge)
	sort.Sort(ByAge(kids))
	fmt.Println("Sort by age", kids)
}
