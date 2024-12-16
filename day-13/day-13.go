package main

import (
	"regexp"
	"strings"
	"utils"
)

func puzzle1(input string) int {
	return solveForMachines(parseMachines(input))
}

func puzzle2(input string) int {
	machines := parseMachines(input)

	for i := range machines {
		machines[i].prize.x += 10000000000000
		machines[i].prize.y += 10000000000000
	}

	return solveForMachines(machines)
}

func solveForMachines(machines []Machine) (result int) {
	for _, machine := range machines {
		diff := machine.buttons[0].x*machine.buttons[1].y - machine.buttons[1].x*machine.buttons[0].y
		prizeDiff := machine.prize.y*machine.buttons[0].x - machine.prize.x*machine.buttons[0].y

		if prizeDiff%diff != 0 {
			continue
		}

		b := prizeDiff / diff
		a := (machine.prize.x - machine.buttons[1].x*b) / machine.buttons[0].x

		result += (a * 3) + b
	}
	return
}

type Machine struct {
	buttons []Button
	prize   Prize
}

type Button XY

type Prize XY

type XY struct {
	x, y int
}

func parseMachines(input string) (machines []Machine) {
	for _, machineString := range strings.Split(input, "\n\n") {
		lines := strings.Split(machineString, "\n")

		machines = append(machines, Machine{
			buttons: []Button{Button(parseXy(lines[0])), Button(parseXy(lines[1]))},
			prize:   Prize(parseXy(lines[2])),
		})
	}
	return
}

func parseXy(input string) XY {
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(input, -1)

	return XY{
		x: utils.ToInt(matches[0]),
		y: utils.ToInt(matches[1]),
	}
}
