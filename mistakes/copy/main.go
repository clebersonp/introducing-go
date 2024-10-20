package main

import "fmt"

func main() {

	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println("dst:", dst) // dst: []

	// the number of elements copied by the copy function is the minimum of len(dst) and len(scr).
	// To make a full copy, you must allocate a big enough destination slice.

	var src2, dst2 []int
	src2 = []int{1, 2, 3}
	dst2 = make([]int, len(src2))
	n := copy(dst2, src2) // return the number of elements copied.
	fmt.Println("dts2:", dst2, "(copied", n, "numbers)")

	// you can also use the append function to make a copy by appending to a nil slice.
	var src3, dst3 []int
	src3 = []int{1, 2, 3}
	dst3 = append(dst3, src3...)
	fmt.Println("dst3:", dst3, "len:", len(dst3), "cap:", cap(dst3))

}
