package main

import (
	"strings"
	"utils"
)

func puzzle1(input string) (result int) {
	m, starts, ends := parseMap(input)

	for _, start := range starts {
		for _, end := range ends {
			if pathExists(start, end, m) {
				result++
			}
		}
	}

	return
}

func puzzle2(input string) (result int) {
	m, starts, ends := parseMap(input)

	for _, start := range starts {
		for _, end := range ends {
			result += countPaths(start, end, m)
		}
	}

	return
}

func pathExists(start Coord, end Coord, m Map) bool {
	return countPaths(start, end, m) > 0
}

func countPaths(start Coord, end Coord, m Map) (result int) {
	for _, c := range []Coord{
		{start.x - 1, start.y},
		{start.x + 1, start.y},
		{start.x, start.y - 1},
		{start.x, start.y + 1},
	} {
		if m[c]-m[start] != 1 {
			continue
		}

		if c == end {
			return 1
		}

		result += countPaths(c, end, m)
	}

	return
}

type Map map[Coord]int

type Coord struct {
	x int
	y int
}

func parseMap(input string) (m Map, starts []Coord, ends []Coord) {
	lines := strings.Split(input, "\n")

	m = make(Map)

	for i, line := range lines {
		parts := strings.Split(line, "")
		for j, part := range parts {
			height := utils.ToInt(part)
			coord := Coord{j, i}

			m[coord] = height

			if height == 0 {
				starts = append(starts, coord)
			}

			if height == 9 {
				ends = append(ends, coord)
			}

		}
	}

	return
}
