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
	x, y := 0, 0

	for i := 1; i < len(sequence)-1; i++ {
		if sequence[i] <= sequence[i-1] {
			x++
		}
		if sequence[i+1] <= sequence[i-1] {
			y++
		}
	}

	if sequence[len(sequence)-1] <= sequence[len(sequence)-2] {
		x++
	}

	if x <= 1 && y <= 1 {
		return true
	}

	return false
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
