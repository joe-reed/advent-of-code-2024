package main

import (
	"slices"
	"strings"

	. "utils"
)

func puzzle1(input string) (result int) {
	rules, updates := parseRulesAndUpdates(input)

	correctUpdates, _ := partitionUpdatesByCorrectness(updates, rules)

	for _, update := range correctUpdates {
		result += ToInt(update.middle())
	}

	return
}

func puzzle2(input string) (result int) {
	rules, updates := parseRulesAndUpdates(input)

	_, incorrectUpdates := partitionUpdatesByCorrectness(updates, rules)

	for _, update := range incorrectUpdates {
		result += ToInt(update.correct(rules).middle())
	}

	return
}

type Rule struct {
	before string
	after  string
}

type Update []string

func (u Update) satisfies(rules []Rule) bool {
	return len(u.getUnsatisfiedRules(rules)) == 0
}

func (u Update) getUnsatisfiedRules(rules []Rule) []Rule {
	var unsatisfiedRules []Rule

	for _, rule := range rules {
		beforeIndex := slices.Index(u, rule.before)
		afterIndex := slices.Index(u, rule.after)

		if beforeIndex == -1 || afterIndex == -1 {
			continue
		}

		if beforeIndex > afterIndex {
			unsatisfiedRules = append(unsatisfiedRules, rule)
		}
	}

	return unsatisfiedRules
}

func (u Update) correct(rules []Rule) Update {
	correctedUpdate := make(Update, len(u))

	copy(correctedUpdate, u)

	unsatisfiedRules := u.getUnsatisfiedRules(rules)

	for len(unsatisfiedRules) > 0 {
		rule := unsatisfiedRules[0]

		beforeIndex := slices.Index(correctedUpdate, rule.before)
		afterIndex := slices.Index(correctedUpdate, rule.after)

		correctedUpdate[beforeIndex], correctedUpdate[afterIndex] = correctedUpdate[afterIndex], correctedUpdate[beforeIndex]

		unsatisfiedRules = correctedUpdate.getUnsatisfiedRules(rules)
	}

	return correctedUpdate
}

func (u Update) middle() string {
	middleIndex := len(u) / 2

	return u[middleIndex]
}

func partitionUpdatesByCorrectness(updates []Update, rules []Rule) ([]Update, []Update) {
	var correctUpdates, incorrectUpdates []Update

	for _, update := range updates {
		if update.satisfies(rules) {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	return correctUpdates, incorrectUpdates
}

func parseRulesAndUpdates(input string) ([]Rule, []Update) {
	parts := strings.Split(input, "\n\n")

	rulesString := parts[0]
	updatesString := parts[1]

	var rules []Rule

	for _, line := range strings.Split(rulesString, "\n") {
		parts := strings.Split(line, "|")
		rules = append(rules, Rule{
			before: parts[0],
			after:  parts[1],
		})
	}

	var updates []Update

	for _, line := range strings.Split(updatesString, "\n") {
		updates = append(updates, strings.Split(line, ","))
	}

	return rules, updates
}
