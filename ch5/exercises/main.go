package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
		4, 14, 7, 3,
	}

	fmt.Println("Numbers:", x)
	small := smallest(x)
	fmt.Println("Smallest number is:", small)
}

func smallest(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	small := numbers[0]
	for _, num := range numbers {
		if num < small {
			small = num
		}
	}
	return small
}
