package main

import (
	"fmt"
	"testing"
	. "utils"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			140,
		},
		{
			"./test-input-2.txt",
			772,
		},
		{
			"./test-input-3.txt",
			1930,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle1(FileToString(test.input)))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(FileToString("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			80,
		},
		{
			"./test-input-2.txt",
			436,
		},
		{
			"./test-input-3.txt",
			1206,
		},
		{
			"./test-input-4.txt",
			236,
		},
		{
			"./test-input-5.txt",
			368,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(FileToString(test.input)))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(FileToString("./input.txt")))
}
