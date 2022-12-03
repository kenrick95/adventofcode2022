package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("day03.in")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\n")

	// Part 1
	var totalPriorities = 0
	for _, line := range lines {
		var cleanedLine string = strings.TrimSpace(line)
		if cleanedLine == "" {
			continue
		}
		// split string as two
		var runes = []rune(cleanedLine)
		var totalLen = len(runes)
		var len = totalLen / 2
		var occurence = map[rune]int{}
		for i := 0; i < len; i++ {
			occurence[runes[i]] = 1
		}

		// find the common character
		var commonCharRune rune
		for i := len; i < totalLen; i++ {
			if occurence[runes[i]] == 1 {
				commonCharRune = runes[i]
				break
			}
		}
		// var commonChar = string(commonCharRune)
		// println("common", commonChar)
		// determine the priority
		var priority int
		if []rune("A")[0] <= commonCharRune &&
			commonCharRune <= []rune("Z")[0] {
			priority = int(commonCharRune) - int(([]rune("A")[0])) + 27
		} else {
			priority = int(commonCharRune) - int(([]rune("a")[0])) + 1
		}
		// sum it up
		totalPriorities += priority
	}
	println("totalPriorities", totalPriorities)

	// Part 2
	var totalPriorities2 = 0
	var occurence = map[rune]int{}

	for lineIndex, line := range lines {
		var cleanedLine string = strings.TrimSpace(line)
		if cleanedLine == "" {
			continue
		}
		// println(lineIndex)
		if lineIndex%3 == 0 {
			// reset occurence map
			occurence = map[rune]int{}
		}

		var runes = []rune(cleanedLine)
		var totalLen = len(runes)
		var occurenceWithinLine = map[rune]int{}
		for i := 0; i < totalLen; i++ {
			occurenceWithinLine[runes[i]] = 1
		}

		for charRune, _ := range occurenceWithinLine {
			currentValue, exists := occurence[charRune]
			if !exists {
				occurence[charRune] = 0
			}
			occurence[charRune] = currentValue + 1
		}

		if lineIndex%3 == 2 {

			// find the common character
			var commonCharRune rune
			for charRune, value := range occurence {
				if value == 3 {
					commonCharRune = charRune
					break
				}
			}
			// var commonChar = string(commonCharRune)
			// println("common", commonChar)
			// determine the priority
			var priority int
			if []rune("A")[0] <= commonCharRune &&
				commonCharRune <= []rune("Z")[0] {
				priority = int(commonCharRune) - int(([]rune("A")[0])) + 27
			} else {
				priority = int(commonCharRune) - int(([]rune("a")[0])) + 1
			}
			// sum it up
			totalPriorities2 += priority
		}
	}
	println("totalPriorities2", totalPriorities2)
}
