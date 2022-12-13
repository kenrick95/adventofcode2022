package main

import (
	"encoding/json"
	"sort"
	"utils"
)

func main() {
	lines := utils.ReadFileToLines("day13.in")
	// Gave up using Go for Day 13... since I need to know the structure of the array first to be able to access the value (dead)
	// OK now reading the solutions, seems possible, let's practice it. Inspiration from  https://www.reddit.com/r/adventofcode/comments/zkmyh4/comment/j01h74o/?context=3

	packetListsPart1 := []any{}
	packetListsPart2 := []any{}
	packetCount := 0
	ansPart1 := 0

	for _, line := range lines {
		if line == "" && len(packetListsPart1) == 2 {
			packetCount += 1
			res := compareFn(packetListsPart1[0], packetListsPart1[1])
			// println("Res", packetCount, res)
			if res == -1 {
				ansPart1 += packetCount
			}

			packetListsPart1 = []any{}
		} else if line == "" {
			continue
		} else {

			var packetList any
			json.Unmarshal([]byte(line), &packetList)
			packetListsPart1 = append(packetListsPart1, packetList)
			packetListsPart2 = append(packetListsPart2, packetList)
		}
	}

	println("Ans Part 1:", ansPart1)

	var divider1 any
	json.Unmarshal([]byte("[[2]]"), &divider1)
	packetListsPart2 = append(packetListsPart2, divider1)
	var divider2 any
	json.Unmarshal([]byte("[[6]]"), &divider2)
	packetListsPart2 = append(packetListsPart2, divider2)

	sort.Slice(packetListsPart2, func(i int, j int) bool {
		return compareFn(packetListsPart2[i], packetListsPart2[j]) <= 0
	})

	ansPart2 := 1

	for i, packet := range packetListsPart2 {
		packetBytes, _ := json.Marshal(packet)
		if string(packetBytes) == "[[2]]" || string(packetBytes) == "[[6]]" {
			ansPart2 *= (i + 1)
		}
	}
	println("Ans Part 2:", ansPart2)
}
func compareFn(packetListA any, packetListB any) int {
	numA, okA := packetListA.(float64)
	numB, okB := packetListB.(float64)
	if okA && okB {
		// Both are number
		if int(numA) < int(numB) {
			return -1
		} else if int(numA) > int(numB) {
			return 1
		}
		return 0

	} else if !okA && !okB {
		// Both are list
		arrA := packetListA.([]any)
		arrB := packetListB.([]any)
		maxLen := 0
		if len(arrA) > len(arrB) {
			maxLen = len(arrA)
		} else {
			maxLen = len(arrB)
		}
		for i := 0; i < maxLen; i++ {
			if i >= len(arrA) {
				return -1
			} else if i >= len(arrB) {
				return 1
			}
			res := compareFn(arrA[i], arrB[i])
			if res != 0 {
				return res
			}
		}
		return 0
	} else if okA && !okB {
		// A is number, B is List
		arrA := []any{packetListA}
		arrB := packetListB.([]any)
		return compareFn(arrA, arrB)
	} else if !okA && okB {
		// A is list, B is number
		arrA := packetListA.([]any)
		arrB := []any{packetListB}
		return compareFn(arrA, arrB)
	}

	return 0
}
