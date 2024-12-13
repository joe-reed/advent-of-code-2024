package main

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
)

func puzzle1(input string) (result int) {
	regions := parseRegions(input)

	return lo.Sum(lo.Map(regions, func(region Region, _ int) int {
		return region.price()
	}))
}

func puzzle2(input string) (result int) {
	regions := parseRegions(input)

	return lo.Sum(lo.Map(regions, func(region Region, _ int) int {
		return region.priceWithDiscount()
	}))
}

type Region []Plot

func (region Region) priceWithDiscount() int {
	return region.numberOfSides() * region.area()
}

func (region Region) price() int {
	return region.perimeter() * region.area()
}

func (region Region) perimeter() int {
	return lo.Sum(lo.Map(region, func(plot Plot, _ int) int {
		adjacentPlotCount := 0

		for _, coord := range [][]int{
			{plot.x, plot.y - 1},
			{plot.x, plot.y + 1},
			{plot.x - 1, plot.y},
			{plot.x + 1, plot.y},
		} {
			_, found := lo.Find(region, func(plot Plot) bool {
				return plot.x == coord[0] && plot.y == coord[1]
			})

			if found {
				adjacentPlotCount++
			}
		}

		return 4 - adjacentPlotCount
	}))
}

type Side struct {
	x         int
	y         int
	direction string
	length    int
}

func (s Side) split(point []int) []Side {
	if s.direction == "horizontal" {
		return []Side{
			{x: s.x, y: s.y, direction: "horizontal", length: point[0] - s.x},
			{x: point[0], y: s.y, direction: "horizontal", length: s.x + s.length - point[0]},
		}
	}

	return []Side{
		{x: s.x, y: s.y, direction: "vertical", length: point[1] - s.y},
		{x: s.x, y: point[1], direction: "vertical", length: s.y + s.length - point[1]},
	}
}

func (s Side) merge(b Side) (Side, error) {
	result := s

	if s.direction != b.direction {
		return result, fmt.Errorf("cannot merge sides with different directions")
	}

	if s.direction == "horizontal" {
		if s.y != b.y {
			return result, fmt.Errorf("cannot merge horizontal sides with different y coordinates")
		}

		if b.x < s.x {
			if b.x+b.length != s.x {
				return result, fmt.Errorf("cannot merge horizontal sides with gap in x coordinates")
			}

			result.x = b.x
		} else {
			if s.x+s.length != b.x {
				return result, fmt.Errorf("cannot merge horizontal sides with gap in x coordinates")
			}

			result.x = s.x
		}

	}

	if s.direction == "vertical" {
		if s.x != b.x {
			return result, fmt.Errorf("cannot merge vertical sides with different x coordinates")
		}

		if b.y < s.y {
			if b.y+b.length != s.y {
				return result, fmt.Errorf("cannot merge vertical sides with gap in y coordinates")
			}

			result.y = b.y
		} else {
			if s.y+s.length != b.y {
				return result, fmt.Errorf("cannot merge vertical sides with gap in y coordinates")
			}

			result.y = s.y
		}
	}

	result.length = s.length + b.length

	return result, nil
}

func (s Side) intersects(b Side) bool {
	if s.direction == b.direction {
		return false
	}

	if s.direction == "horizontal" {
		if s.y > b.y && s.y < b.y+b.length && b.x > s.x && b.x < s.x+s.length {
			return true
		}
	}

	if s.direction == "vertical" {
		if s.x > b.x && s.x < b.x+b.length && b.y > s.y && b.y < s.y+s.length {
			return true
		}
	}

	return false
}

func (region Region) numberOfSides() int {
	sides := []Side{}

	for _, plot := range region {
		_, found := lo.Find(region, func(p Plot) bool {
			return p.x == plot.x && p.y == plot.y-1
		})

		if !found {
			sides = append(sides, Side{x: plot.x, y: plot.y, direction: "horizontal", length: 1})
		}

		_, found = lo.Find(region, func(p Plot) bool {
			return p.x == plot.x && p.y == plot.y+1
		})

		if !found {
			sides = append(sides, Side{x: plot.x, y: plot.y + 1, direction: "horizontal", length: 1})
		}

		_, found = lo.Find(region, func(p Plot) bool {
			return p.x == plot.x-1 && p.y == plot.y
		})

		if !found {
			sides = append(sides, Side{x: plot.x, y: plot.y, direction: "vertical", length: 1})
		}

		_, found = lo.Find(region, func(p Plot) bool {
			return p.x == plot.x+1 && p.y == plot.y
		})

		if !found {
			sides = append(sides, Side{x: plot.x + 1, y: plot.y, direction: "vertical", length: 1})
		}
	}

	mergedSides := mergeSides(sides)

	intersectionCount := 0
	for i := 0; i < len(mergedSides); i++ {
		for j := 0; j < len(mergedSides); j++ {
			if mergedSides[i].intersects(mergedSides[j]) {
				intersectionCount++
			}
		}
	}

	return len(mergedSides) + intersectionCount
}

func mergeSides(sides []Side) (mergedSides []Side) {
	for i := 0; i < len(sides); i++ {
		canMerge := false
		for j := 0; j < len(mergedSides); j++ {
			merged, err := mergedSides[j].merge(sides[i])

			if err != nil {
				continue
			}

			canMerge = true
			mergedSides = slices.Delete(mergedSides, j, j+1)
			mergedSides = append(mergedSides, merged)
			break
		}

		if !canMerge {
			mergedSides = append(mergedSides, sides[i])
		}
	}

	if len(mergedSides) == len(sides) {
		return mergedSides
	}

	return mergeSides(mergedSides)
}

func (region Region) area() int {
	return len(region)
}

type Plot struct {
	x     int
	y     int
	plant string
}

func parseRegions(input string) (regions []Region) {
	lines := strings.Split(input, "\n")

	var plots []Plot

	for i, line := range lines {
		chars := strings.Split(line, "")

		for j, char := range chars {
			plots = append(plots, Plot{x: j, y: i, plant: char})
		}
	}

	for len(plots) > 0 {
		region := Region{}

		newlyAdded := []Plot{plots[0]}

		for len(newlyAdded) > 0 {
			plots = slices.DeleteFunc(plots, func(plot Plot) bool {
				return lo.Contains(newlyAdded, plot)
			})

			for _, plot := range newlyAdded {
				region = append(region, plot)
			}

			newlyAdded = getAdjacentPlots(region, newlyAdded, plots)
		}

		regions = append(regions, region)
	}

	return
}

func getAdjacentPlots(region Region, newlyAdded []Plot, plots []Plot) []Plot {
	adjacentPlots := []Plot{}

	for _, plot := range newlyAdded {
		adjacentForPlot := []Plot{}
		for _, coord := range [][]int{
			{plot.x, plot.y - 1},
			{plot.x, plot.y + 1},
			{plot.x - 1, plot.y},
			{plot.x + 1, plot.y},
		} {
			plot, found := lo.Find(plots, func(plot Plot) bool {
				return plot.x == coord[0] && plot.y == coord[1]
			})

			if found {
				adjacentForPlot = append(adjacentForPlot, plot)
			}
		}

		newPlotsForRegion := lo.Filter(adjacentForPlot, func(plot Plot, _ int) bool {
			return plot.plant == newlyAdded[0].plant && !lo.Contains(region, plot) && !lo.Contains(adjacentPlots, plot)
		})

		for _, plot := range newPlotsForRegion {
			adjacentPlots = append(adjacentPlots, plot)
		}
	}

	return adjacentPlots
}
