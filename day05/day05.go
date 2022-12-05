package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

func main() {
	lines := utils.ReadFileToLines("day05.in")

	// Parsing input, split between initial states & commands
	var initialStateLines []string
	var commandLines []string
	var hasSep = false
	for _, line := range lines {
		if line == "" {
			if hasSep {
				break
			} else {
				hasSep = true
			}
		} else {
			if hasSep {
				commandLines = append(commandLines, line)
			} else {
				initialStateLines = append(initialStateLines, line)
			}
		}
	}

	// gosh parsing it is really crazy, might as well hardcode it 3:)
	// initialState := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }
	initialState := [][]string{
		{"N", "B", "D", "T", "V", "G", "Z", "J"},
		{"S", "R", "M", "D", "W", "P", "F"},
		{"V", "C", "R", "S", "Z"},
		{"R", "T", "J", "Z", "P", "H", "G"},
		{"T", "C", "J", "N", "D", "Z", "Q", "F"},
		{"N", "V", "P", "W", "G", "S", "F", "M"},
		{"G", "C", "V", "B", "P", "Q"},
		{"Z", "B", "P", "N"},
		{"W", "P", "J"},
	}
	statePart1 := make([][]string, len(initialState))
	copy(statePart1, initialState)
	statePart2 := make([][]string, len(initialState))
	copy(statePart2, initialState)

	re, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)

	for _, commandLine := range commandLines {
		matches := re.FindStringSubmatch(commandLine)
		commandCount, _ := strconv.Atoi(matches[1])
		commandSource, _ := strconv.Atoi(matches[2])
		commandDestination, _ := strconv.Atoi(matches[3])
		commandSource -= 1
		commandDestination -= 1

		// https://github.com/golang/go/wiki/SliceTricks

		// println("cmd", commandLine, commandCount, commandSource, commandDestination)

		var elementsHoldPart1 []string
		var elementsHoldPart2 []string
		for i := 0; i < commandCount; i++ {
			// Pop from state (source)
			var elementPart1 string
			var elementPart2 string

			elementPart1, statePart1[commandSource] = statePart1[commandSource][len(statePart1[commandSource])-1], statePart1[commandSource][:len(statePart1[commandSource])-1]

			elementPart2, statePart2[commandSource] = statePart2[commandSource][len(statePart2[commandSource])-1], statePart2[commandSource][:len(statePart2[commandSource])-1]

			// Push to elementsHold
			elementsHoldPart1 = append(elementsHoldPart1, elementPart1)
			elementsHoldPart2 = append([]string{elementPart2}, elementsHoldPart2...)
		}
		for i := 0; i < commandCount; i++ {
			// Push to state (destination)
			statePart1[commandDestination] = append(statePart1[commandDestination], elementsHoldPart1[i])

			statePart2[commandDestination] = append(statePart2[commandDestination], elementsHoldPart2[i])
		}
		fmt.Printf("%v\n", statePart1)
		fmt.Printf("%v\n", statePart2)

	}

	var ansPart1 []string
	for _, stack := range statePart1 {
		ansPart1 = append(ansPart1, stack[len(stack)-1])
	}
	var ansPart2 []string
	for _, stack := range statePart2 {
		ansPart2 = append(ansPart2, stack[len(stack)-1])
	}

	fmt.Println("Part 1:", strings.Join(ansPart1, ""))
	fmt.Println("Part 2:", strings.Join(ansPart2, ""))

}
