package main

import (
	"slices"
	"strconv"
	"strings"
	"utils"
)

func puzzle1(input string) (result int) {
	return solve(input, 25)
}

//func puzzle2(input string) (result int) {
//	return solve(input, 75)
//}

func solve(input string, blinkNumber int) (result int) {
	stones := parseStones(input)

	for range blinkNumber {
		stones = stones.blink()
	}

	return len(stones)
}

type Stones []int

func (stones Stones) blink() Stones {
	newStones := make(Stones, len(stones))

	j := 0
	for i := 0; i < len(stones); i++ {
		if stones[i] == 0 {
			newStones[j] = 1

			j++
			continue
		}

		stoneString := strconv.Itoa(stones[i])
		if len(stoneString)%2 == 0 {
			halfway := len(stoneString) / 2

			newStones[j] = utils.ToInt(stoneString[:halfway])
			newStones = slices.Insert(newStones, j+1, utils.ToInt(stoneString[halfway:]))

			j += 2

			continue
		}

		newStones[j] = stones[i] * 2024
		j++
		continue
	}

	return newStones
}

func parseStones(input string) (stones Stones) {
	return utils.MapToInts(strings.Split(input, " "))
}
