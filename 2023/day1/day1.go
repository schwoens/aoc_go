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
	fmt.Printf("Part 2: %d\n", part2(input))
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

var writtenDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func getDigit(writtenDigit string) rune {
	switch writtenDigit {
	case "one":
		return '1'
	case "two":
		return '2'
	case "three":
		return '3'
	case "four":
		return '4'
	case "five":
		return '5'
	case "six":
		return '6'
	case "seven":
		return '7'
	case "eight":
		return '8'
	case "nine":
		return '9'
	default:
		panic("Invalid digit")
	}
}

func part2(input string) int {
	trimmedInput := strings.TrimSuffix(input, "\n")
	sum := 0

	for _, line := range strings.Split(trimmedInput, "\n") {
		var firstDigit rune
		var lastDigit rune

		firstDigitIndex := len(line)
		lastDigitIndex := -1

		for _, writtenDigit := range writtenDigits {
			index := strings.Index(line, writtenDigit)
			if index != -1 && index < firstDigitIndex {
				firstDigitIndex = index
				firstDigit = getDigit(writtenDigit)
			}

			index = strings.LastIndex(line, writtenDigit)
			if index != -1 && index > lastDigitIndex {
				lastDigitIndex = index
				lastDigit = getDigit(writtenDigit)
			}
		}

		for i, char := range line {
			if unicode.IsDigit(char) {
				if i < firstDigitIndex {
					firstDigit = char
				}
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				if i > lastDigitIndex {
					lastDigit = char
				}
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
