package main

import (
	"github.com/samber/lo"
	"maps"
	"strconv"
	"strings"
	"utils"
)

func puzzle1(input string) (result int) {
	return solve(input, 25)
}

func puzzle2(input string) (result int) {
	return solve(input, 75)
}

func solve(input string, blinkNumber int) (result int) {
	counts := parseCounts(input)

	for range blinkNumber {
		newCounts := make(map[Stone]int)

		for stone := range counts {
			for _, s := range stone.blink() {
				newCounts[s] += counts[stone]
			}
		}

		counts = newCounts
	}

	for count := range maps.Values(counts) {
		result += count
	}

	return
}

type Stone int

func (stone Stone) blink() []Stone {
	if stone == 0 {
		return []Stone{1}
	}

	stoneString := strconv.Itoa(int(stone))

	if len(stoneString)%2 == 0 {
		halfway := len(stoneString) / 2

		return []Stone{
			Stone(utils.ToInt(stoneString[:halfway])),
			Stone(utils.ToInt(stoneString[halfway:])),
		}
	}

	return []Stone{stone * 2024}
}

func parseCounts(input string) (counts map[Stone]int) {
	stones := lo.Map(strings.Split(input, " "), func(s string, _ int) Stone {
		return Stone(utils.ToInt(s))
	})

	counts = make(map[Stone]int)
	for _, stone := range stones {
		counts[stone] = 1
	}

	return
}
