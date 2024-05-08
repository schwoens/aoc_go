package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("input could not be read")
	}
	input := string(bytes)
	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input string) int {
	trimmedInput := strings.TrimSuffix(input, "\n")
	sum := 0

	for _, line := range strings.Split(trimmedInput, "\n") {
		var firstDigit rune
		var lastDigit rune

		for _, char := range line {
			if unicode.IsDigit(char) {
				firstDigit = char
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				lastDigit = char
				break
			}
		}

		numString := string([]rune{firstDigit, lastDigit})
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic("error parsing num to int")
		}
		sum += num
	}
	return sum
}
