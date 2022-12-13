package main

import (
	"utils"
	// "github.com/buger/jsonparser"
)

func main() {
	lines := utils.ReadFileToLines("day13.in")
	// Gave up using Go for Day 13... since I need to know the structure of the array first to be able to access the value (dead)

	packetLists := []string{}
	packetCount := 0
	ansPart1 := 0
	for _, line := range lines {
		if line == "" && len(packetLists) == 2 {
			packetCount += 1
			if isCorrectOrder(packetLists[0], packetLists[1]) {
				ansPart1 += packetCount
			}

			packetLists = []string{}
		}

		packetLists = append(packetLists, line)
	}

	println("Ans Part 1:", ansPart1)
}
func isCorrectOrder(packetListA string, packetListB string) bool {
	return true
}

// func parseLine(line string) any {
// 	if strings.HasPrefix(line, "[") {
// 		// List
// 		chars := strings.Split(line, "")
// 		subStrings := []string{}
// 		level := 0
// 		curChars := []string{}
// 		for i := 1; i < len(chars)-1; i++ {
// 			if level == 0 && chars[i] == "," {
// 				subStrings = append(subStrings, strings.Join(curChars, ""))
// 			} else {
// 				if chars[i] == "[" {
// 					level += 1
// 				} else if chars[i] == "]" {
// 					level -= 1
// 				}
// 				curChars = append(curChars, chars[i])
// 			}
// 		}

// 	} else {
// 		// Integer

// 	}
// }
