package main

import "fmt"

// enum example from: https://yourbasic.org/golang/iota/

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Value() int {
	return int(d)
}

func (d Direction) Name() string {
	return d.String()
}
func (d Direction) String() string {
	return []string{North: "North", East: "East", South: "South", West: "West"}[d]
}

func main() {
	var d = North
	fmt.Println(d.Name(), ":", d.Value())
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	case East:
		fmt.Println(" goes right.")
	case West:
		fmt.Println(" goes left.")
	default:
		fmt.Println(" stays put.")
	}

	// Output: North goes up.
}
