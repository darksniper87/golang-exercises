package intro

func add(param1 int, param2 int) int {
	return param1 + param2
}

func centuryFromYear(year int) int {
	return (year + 99) / 100
}

func checkPalindrome(inputString string) bool {
	for i, j := 0, len(inputString)-1; i < j; {
		if inputString[i] != inputString[j] {
			return false
		}
		i++
		j--
	}

	return true
}
