package main

import (
	"strconv"
	"strings"
	"utils"
)

func main() {
	cycleNumber := 1

	// signal strength (the cycle number multiplied by the value of the X register)
	// Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles.
	ansPart1 := 0
	ansPart2 := []string{}
	// sprite position is ${registerX - 1} till ${registerX + 1}
	registerX := 1
	lines := utils.ReadFileToLines("day10.in")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "noop" {
			ansPart1 += getSignalStrength(cycleNumber, registerX)
			ansPart2 = append(ansPart2, getPixel(cycleNumber, registerX))
			cycleNumber += 1
		} else if parts[0] == "addx" {
			amount, _ := strconv.Atoi(parts[1])
			ansPart1 += getSignalStrength(cycleNumber, registerX)
			ansPart2 = append(ansPart2, getPixel(cycleNumber, registerX))
			cycleNumber += 1
			ansPart1 += getSignalStrength(cycleNumber, registerX)
			ansPart2 = append(ansPart2, getPixel(cycleNumber, registerX))
			cycleNumber += 1
			registerX += amount
		}
	}
	println("Part 1:", ansPart1)

	println("Part 2:")
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			print(ansPart2[i*40+j])
		}
		println()
	}
}

func getSignalStrength(cycleNumber int, registerX int) int {
	if cycleNumber != 20 && cycleNumber != 60 && cycleNumber != 100 && cycleNumber != 140 && cycleNumber != 180 && cycleNumber != 220 {
		return 0
	}
	signalStrength := cycleNumber * registerX
	return signalStrength
}

func getPixel(cycleNumber int, registerX int) string {
	pixelPosition := (cycleNumber % 40) - 1
	if registerX-1 <= pixelPosition && pixelPosition <= registerX+1 {
		return "#"
	}
	return "."
}
