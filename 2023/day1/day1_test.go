package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("test_input.txt")
	if err != nil {
		panic("test input could not be read")
	}

	input := string(bytes)
	result := part1(input)

	if result != 142 {
		t.Errorf("Part 1: %d; want 142", result)
	}
}
