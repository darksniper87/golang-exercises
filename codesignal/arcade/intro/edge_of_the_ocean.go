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
