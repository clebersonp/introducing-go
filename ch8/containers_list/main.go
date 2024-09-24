package main

import (
	"container/list"
	"fmt"
)

func main() {

	// Doubly linked list
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	// forward
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
	fmt.Println()

	// backward
	for e := x.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}

	// A *List can also be created using list.New
	//l := list.New()

	//
}
