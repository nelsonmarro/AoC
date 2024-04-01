package txthelpers

import (
	"fmt"
	"regexp"
	"strconv"
)

func ExtractNumbers(text string) []int {
	var numbers []int
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(text, -1)

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func ExtractColors(text string) []string {
	re := regexp.MustCompile("[a-zA-Z]+")
	matches := re.FindAllString(text, -1)
	return matches
}
