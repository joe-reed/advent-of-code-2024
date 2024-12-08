package main

import (
	"strconv"
	"strings"

	. "utils"
)

func puzzle1(input string) (result int) {
	return solvePuzzleForRules(input, []Rule{divisionRule, subtractionRule})
}

func puzzle2(input string) (result int) {
	return solvePuzzleForRules(input, []Rule{divisionRule, subtractionRule, concatenationRule})
}

func solvePuzzleForRules(input string, rules []Rule) (result int) {
	for _, equation := range parseEquations(input) {
		if equation.isSolvable(rules) {
			result += equation.result
		}
	}
	return
}

type Equation struct {
	result   int
	operands []int
}

func (e Equation) isSolvable(rules []Rule) bool {
	i := len(e.operands) - 1
	operand := e.operands[i]

	if i == 0 {
		return e.result == operand
	}

	for _, rule := range rules {
		if rule.test(e.result, operand) && (Equation{rule.getNewResult(e.result, operand), e.operands[:i]}).isSolvable(rules) {
			return true
		}
	}

	return false
}

type Rule struct {
	test         func(int, int) bool
	getNewResult func(int, int) int
}

var divisionRule = Rule{
	func(result, operand int) bool {
		return result%operand == 0
	},
	func(result, operand int) int {
		return result / operand
	},
}

var concatenationRule = Rule{
	func(result, operand int) bool {
		newResult, hasSuffix := strings.CutSuffix(strconv.Itoa(result), strconv.Itoa(operand))

		return hasSuffix && newResult != ""
	},
	func(result, operand int) int {
		newResult, _ := strings.CutSuffix(strconv.Itoa(result), strconv.Itoa(operand))

		return ToInt(newResult)
	},
}

var subtractionRule = Rule{
	func(result, operand int) bool {
		return result > operand
	},
	func(result, operand int) int {
		return result - operand
	},
}

func parseEquations(input string) (equations []Equation) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		equation := Equation{}
		parts := strings.Split(line, ": ")

		equation.result = ToInt(parts[0])
		equation.operands = MapToInts(strings.Split(parts[1], " "))

		equations = append(equations, equation)
	}

	return
}
