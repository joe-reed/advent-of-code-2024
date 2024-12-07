package main

import (
	"strings"
)

func puzzle1(input string) (result int) {
	m := parseMap(input)

	for m.guardInBounds() {
		m.moveGuard()
	}

	return len(m.visited) - 1
}

func puzzle2(input string) (result int) {
	m := parseMap(input)

	for m.guardInBounds() {
		m.moveGuard()
	}

	visited := m.visited

	for _, empty := range visited {
		if empty == m.startingPosition {
			continue
		}

		m.reset()

		m.addObstruction(empty)

		pathLength := 1
		for m.guardInBounds() {
			m.moveGuard()
			pathLength++

			if pathLength > 6000 {
				result++
				break
			}
		}
	}

	return result
}

type Map struct {
	guard          Coord
	guardDirection string
	size           int
	obstructions   []Coord
	empty          []Coord
	visited        []Coord

	startingPosition     Coord
	startingObstructions []Coord
}

type Coord struct {
	x int
	y int
}

func (m *Map) moveGuard() {
	switch m.guardDirection {
	case "up":
		if m.isObstructed(Coord{m.guard.x, m.guard.y - 1}) {
			m.guardDirection = "right"
			m.moveGuard()
			return
		}
		m.guard = Coord{m.guard.x, m.guard.y - 1}
	case "right":
		if m.isObstructed(Coord{m.guard.x + 1, m.guard.y}) {
			m.guardDirection = "down"
			m.moveGuard()
			return
		}
		m.guard = Coord{m.guard.x + 1, m.guard.y}
	case "down":
		if m.isObstructed(Coord{m.guard.x, m.guard.y + 1}) {
			m.guardDirection = "left"
			m.moveGuard()
			return
		}
		m.guard = Coord{m.guard.x, m.guard.y + 1}
	case "left":
		if m.isObstructed(Coord{m.guard.x - 1, m.guard.y}) {
			m.guardDirection = "up"
			m.moveGuard()
			return
		}
		m.guard = Coord{m.guard.x - 1, m.guard.y}
	}

	if !m.hasVisited(m.guard) {
		m.visited = append(m.visited, m.guard)
	}
}

func (m *Map) isObstructed(c Coord) bool {
	for _, obstruction := range m.obstructions {
		if obstruction == c {
			return true
		}
	}

	return false
}

func (m *Map) hasVisited(c Coord) bool {
	for _, visited := range m.visited {
		if visited == c {
			return true
		}
	}

	return false
}

func (m *Map) guardInBounds() bool {
	return m.guard.x >= 0 && m.guard.y >= 0 && m.guard.x < m.size && m.guard.y < m.size
}

func (m *Map) addObstruction(c Coord) {
	m.obstructions = append(m.obstructions, c)
}

func (m *Map) reset() {
	m.guard = m.startingPosition
	m.guardDirection = "up"
	m.visited = []Coord{m.startingPosition}
	m.obstructions = m.startingObstructions
}

func parseMap(input string) (m Map) {
	lines := strings.Split(input, "\n")

	m.size = len(lines)

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				m.obstructions = append(m.obstructions, Coord{x, y})
			}

			if char == '^' {
				m.guard = Coord{x, y}
			}

			if char == '.' {
				m.empty = append(m.empty, Coord{x, y})
			}
		}
	}

	m.startingPosition = m.guard
	m.startingObstructions = make([]Coord, len(m.obstructions))
	m.startingObstructions = m.obstructions
	m.visited = []Coord{m.startingPosition}
	m.guardDirection = "up"

	return
}
