package main

import "strings"

func puzzle1(input string) (result int) {
	m := parseMap(input)

	antinodes := map[Coord]bool{}
	for _, pair := range m.pairs {

		antinodeA := pair.a.subtractLDistance(pair.lDistance)
		antinodeB := pair.b.addLDistance(pair.lDistance)

		if antinodeA.isWithinBounds(m.size) {
			antinodes[antinodeA] = true
		}

		if antinodeB.isWithinBounds(m.size) {
			antinodes[antinodeB] = true
		}
	}

	return len(antinodes)
}

func puzzle2(input string) (result int) {
	m := parseMap(input)

	antinodes := map[Coord]bool{}
	for _, pair := range m.pairs {
		antinode := pair.a

		for true {
			antinode = antinode.subtractLDistance(pair.lDistance)

			if !antinode.isWithinBounds(m.size) {
				break
			}
		}

		for true {
			antinode = antinode.addLDistance(pair.lDistance)

			if !antinode.isWithinBounds(m.size) {
				break
			}

			antinodes[antinode] = true
		}
	}

	return len(antinodes)
}

type Map struct {
	size  int
	pairs []Pair
}

type Coord struct {
	x int
	y int
}

type Antenna struct {
	frequency string
	coord     Coord
}

type LDistance struct {
	x int
	y int
}

type Pair struct {
	a Coord
	b Coord

	lDistance LDistance
}

func (c Coord) addLDistance(d LDistance) Coord {
	return Coord{c.x + d.x, c.y + d.y}
}

func (c Coord) subtractLDistance(d LDistance) Coord {
	return Coord{c.x - d.x, c.y - d.y}
}

func (c Coord) isWithinBounds(size int) bool {
	return c.x >= 0 && c.x < size && c.y >= 0 && c.y < size
}

func parseMap(input string) (result Map) {
	lines := strings.Split(input, "\n")

	var antennae []Antenna

	for i, line := range lines {
		points := strings.Split(line, "")

		for j, point := range points {
			if point != "." {
				antennae = append(antennae, Antenna{point, Coord{j, i}})
			}
		}
	}

	var pairs []Pair
	for i, a := range antennae {
		for j, b := range antennae {
			if i == j {
				continue
			}

			if a.frequency != b.frequency {
				continue
			}

			pairs = append(pairs, Pair{a.coord, b.coord, LDistance{b.coord.x - a.coord.x, b.coord.y - a.coord.y}})
		}
	}

	return Map{len(lines), pairs}
}
