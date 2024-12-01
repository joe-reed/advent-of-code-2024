package main

import (
	"math"
	"sort"
	"strings"
	. "utils"
)

func puzzle1(input []string) (result int) {
	list1, list2 := getLists(input)

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		result += int(math.Abs(float64(diff)))
	}

	return
}

func puzzle2(input []string) (result int) {
	list1, list2 := getLists(input)

	for i := 0; i < len(list1); i++ {
		count := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				count++
			}
		}
		result += count * list1[i]
		count = 0
	}

	return
}

func getLists(input []string) (list1 []int, list2 []int) {
	for _, line := range input {
		split := strings.Split(line, "   ")

		list1 = append(list1, ToInt(split[0]))
		list2 = append(list2, ToInt(split[1]))
	}

	return
}
