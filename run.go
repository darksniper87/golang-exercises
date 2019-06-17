package main

import (
	"fmt"
)

func main() {
	test := []int{3, 6, 5, 8, 10, 20, 15}
	fmt.Println(almostIncreasingSequence(test))
}

func almostIncreasingSequence(sequence []int) bool {
	errorX, errorY := 0, 0

	for i := 0; i < len(sequence)-2 && errorX < 2 && errorY < 2; i++ {
		if sequence[i] >= sequence[i+1] {
			errorX++
		}

		if sequence[i] >= sequence[i+2] {
			errorY++
		}
	}

	if sequence[len(sequence)-2] >= sequence[len(sequence)-1] {
		errorX++
	}

	return errorX < 2 && errorY < 2
}
