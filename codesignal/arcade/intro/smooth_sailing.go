package intro

import (
	"unicode/utf8"
)

func allLongestStrings(inputArray []string) []string {
    max := 0
	strings := []string{}

	for _, s := range inputArray {
		lenght := utf8.RuneCountInString(s)

		if lenght > max {
			max = lenght
			strings = []string{s}
		} else if lenght == max {
			strings = append(strings, s)
		}
	}

	return strings
}
