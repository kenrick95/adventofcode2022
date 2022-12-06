package main

import (
	"utils"
)

func main() {
	lines := utils.ReadFileToLines("day06.in")
	line := lines[0]
	chars := []rune(line)
	occurrencePart1 := map[rune]int{}
	occurrencePart2 := map[rune]int{}
	ansPart1 := len(chars)
	ansPart2 := len(chars)
	// 0123456789
	// mjqjpqmgbl
	for i, uniqueCount := 0, 0; i < len(chars); i++ {
		// "forget" the past 4 char
		if i >= 4 {
			pastValue := occurrencePart1[chars[i-4]]
			if pastValue == 1 {
				delete(occurrencePart1, chars[i-4])
				uniqueCount -= 1
			} else {
				occurrencePart1[chars[i-4]] = pastValue - 1
			}
		}
		// check current char
		currentValue, exists := occurrencePart1[chars[i]]
		if !exists {
			uniqueCount += 1
			occurrencePart1[chars[i]] = 0
		}
		occurrencePart1[chars[i]] = currentValue + 1
		if uniqueCount == 4 {
			ansPart1 = i + 1
			break
		}
	}

	// Part 2: same as above, but 14
	for i, uniqueCount := 0, 0; i < len(chars); i++ {
		// "forget" the past 14 char
		if i >= 14 {
			pastValue := occurrencePart2[chars[i-14]]
			if pastValue == 1 {
				delete(occurrencePart2, chars[i-14])
				uniqueCount -= 1
			} else {
				occurrencePart2[chars[i-14]] = pastValue - 1
			}
		}
		// check current char
		currentValue, exists := occurrencePart2[chars[i]]
		if !exists {
			uniqueCount += 1
			occurrencePart2[chars[i]] = 0
		}
		occurrencePart2[chars[i]] = currentValue + 1
		if uniqueCount == 14 {
			ansPart2 = i + 1
			break
		}
	}

	println("Part 1:", ansPart1)
	println("Part 2:", ansPart2)
}
