package main

import (
	"strings"
)

func puzzle1(input []string) (result int) {
	allLines := parseAllLines(input)

	for _, line := range allLines {
		result += strings.Count(line, "XMAS")
	}

	return
}

func puzzle2(input []string) (result int) {
	allSquares := parseAllSquares(input)

	for _, square := range allSquares {
		if square[1][1] != "A" {
			continue
		}

		if square[0][0] == "M" && square[0][2] == "M" && square[2][0] == "S" && square[2][2] == "S" {
			result++
		}

		if square[0][0] == "S" && square[0][2] == "S" && square[2][0] == "M" && square[2][2] == "M" {
			result++
		}

		if square[0][0] == "M" && square[0][2] == "S" && square[2][0] == "M" && square[2][2] == "S" {
			result++
		}

		if square[0][0] == "S" && square[0][2] == "M" && square[2][0] == "S" && square[2][2] == "M" {
			result++
		}
	}

	return
}

func parseAllLines(input []string) (result []string) {
	for _, line := range input {
		result = append(result, line)
		result = append(result, reverse(line))
	}

	for i := 0; i < len(input[0]); i++ {
		column := ""
		for j := 0; j < len(input); j++ {
			column += string(input[j][i])
		}

		result = append(result, column)
		result = append(result, reverse(column))
	}

	for i := 0; i < len(input[0]); i++ {
		diagonal := ""

		for j := 0; i+j < len(input); j++ {
			diagonal += string(input[j][i+j])
		}

		result = append(result, diagonal)
		result = append(result, reverse(diagonal))
	}

	for i := 1; i < len(input); i++ {
		diagonal := ""

		for j := 0; j+i < len(input[0]); j++ {
			diagonal += string(input[i+j][j])
		}

		result = append(result, diagonal)
		result = append(result, reverse(diagonal))
	}

	for j := 0; j < len(input[0]); j++ {
		diagonal := ""

		for i := 0; j-i >= 0; i++ {
			diagonal += string(input[i][j-i])
		}

		result = append(result, diagonal)
		result = append(result, reverse(diagonal))
	}

	for i := 1; i < len(input[0]); i++ {
		diagonal := ""

		for j := 0; i+j < len(input); j++ {
			diagonal += string(input[i+j][len(input)-1-j])
		}

		result = append(result, diagonal)
		result = append(result, reverse(diagonal))
	}

	return
}

func parseAllSquares(input []string) (result [][][]string) {
	for i := 0; i <= len(input)-3; i++ {
		for j := 0; j <= len(input[0])-3; j++ {
			square := make([][]string, 3)
			for i := range square {
				square[i] = make([]string, 3)
			}

			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					square[k][l] = string(input[i+k][j+l])
				}
			}

			result = append(result, square)
		}
	}

	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
