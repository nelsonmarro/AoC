package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type numberFound struct {
	textNum     string
	wordNum     string
	appearances []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var totalSum int

	for scanner.Scan() {
		line := scanner.Text()
		totalSum += sumNumbersInTextLine(line)
	}
	file.Close()
	fmt.Printf("Total: %d\n", totalSum)
}

func sumNumbersInTextLine(line string) int {
	line = convertWordsToNumbers(line)
	numbersFromDigit := extractNumberFromDigit(line)

	if len(numbersFromDigit) > 0 {
		firstNum := numbersFromDigit[0]
		lastNum := numbersFromDigit[len(numbersFromDigit)-1]

		// If there's only one number, count it twice.
		if len(numbersFromDigit) == 1 {
			firstNum, lastNum = numbersFromDigit[0], numbersFromDigit[0]
		}

		concatedNums := fmt.Sprint(firstNum, lastNum)
		finalNum, _ := strconv.Atoi(concatedNums)

		return finalNum
	}
	return 0
}
func convertWordsToNumbers(line string) string {
	var numbersFound []numberFound
	numberWordMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for key, value := range numberWordMap {
		if strings.Contains(line, key) {
			numbersFound = append(numbersFound, numberFound{
				textNum:     value,
				appearances: findSubstringIndices(line, key),
			})
		}
	}

	sort.Slice(numbersFound, func(i, j int) bool { return numbersFound[i].appearances[0] < numbersFound[j].appearances[0] })

	for _, numFound := range numbersFound {
		for _, appearance := range numFound.appearances {
			start := line[:appearance+1]
			rest := line[1+appearance:]
			line = fmt.Sprint(start, numFound.textNum, rest)
		}
	}
	return line
}
func extractNumberFromDigit(line string) []string {
	var selectedNumbers []string
	re := regexp.MustCompile("[0-9]")

	matches := re.FindAllString(line, -1)

	for _, match := range matches {
		selectedNumbers = append(selectedNumbers, match)
	}

	return selectedNumbers
}
func findSubstringIndices(s, substr string) []int {
	var indices []int
	start := 0
	for {
		idx := strings.Index(s[start:], substr)
		if idx == -1 {
			return indices
		}
		indices = append(indices, start+idx)
		start = start + idx + 1
	}
}
