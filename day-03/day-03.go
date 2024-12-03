package main

import (
	"regexp"

	. "utils"
)

func puzzle1(input string) (result int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		result += ToInt(match[1]) * ToInt(match[2])
	}

	return
}

func puzzle2(input string) (result int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		}

		if match[0] == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		result += ToInt(match[1]) * ToInt(match[2])
	}

	return
}
