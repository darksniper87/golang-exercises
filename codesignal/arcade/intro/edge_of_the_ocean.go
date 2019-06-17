package intro

import (
	"sort"
)

func adjacentElementsProduct(inputArray []int) int {
	max := inputArray[0] * inputArray[1]

	for i := 2; i < len(inputArray); i++ {
		product := inputArray[i-1] * inputArray[i]

		if product > max {
			max = product
		}
	}

	return max
}

func shapeArea(n int) int {
	if n == 1 {
		return 1
	}

	return n*n + (n-1)*(n-1)
}

func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	length := len(statues)

	return (statues[length-1] - statues[0] + 1) - length
}

func almostIncreasingSequence(sequence []int) bool {
	errorX, errorY := 0, 0

	for i := 0; i < len(sequence)-2; i++ {
		if sequence[i] >= sequence[i+1] {
			errorX++
		}

		if sequence[i] >= sequence[i+2] {
			errorY++
		}

		if errorX > 1 || errorY > 1 {
			return false
		}
	}

	if sequence[len(sequence)-2] >= sequence[len(sequence)-1] && errorX > 0 {
		return false
	}

	return true
}

func matrixElementsSum(matrix [][]int) int {
	sum := 0

	for i := range matrix[0] {
		for j := range matrix {
			if matrix[j][i] == 0 {
				break
			}
			sum += matrix[j][i]
		}
	}

	return sum
}
