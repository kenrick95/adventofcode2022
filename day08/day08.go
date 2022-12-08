package main

import (
	"strconv"
	"strings"
	"utils"
)

type Coord struct {
	x int
	y int
}
type ForestTree struct {
	height      int
	visible     bool
	scenicScore int
}

func main() {
	lines := utils.ReadFileToLines("day08.in")
	forestMap := map[Coord]*ForestTree{}
	var rowCount, columnCount int

	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, char := range strings.Split(line, "") {
			if rowCount == 0 {
				columnCount += 1
			}
			height, _ := strconv.Atoi(char)
			forestMap[Coord{x, y}] = &ForestTree{height: height, visible: false, scenicScore: 0}
		}
		rowCount += 1
	}

	for y := 0; y < rowCount; y++ {
		// Sweep from left
		{
			forestMap[Coord{x: 0, y: y}].visible = true
			currentHeight := forestMap[Coord{x: 0, y: y}].height
			for x := 1; x < columnCount; x++ {
				if forestMap[Coord{x: x, y: y}].height > currentHeight {
					currentHeight = forestMap[Coord{x: x, y: y}].height
					forestMap[Coord{x: x, y: y}].visible = true
				} else {
					continue
				}
			}
		}
		// Sweep from right
		{
			forestMap[Coord{x: columnCount - 1, y: y}].visible = true
			currentHeight := forestMap[Coord{x: columnCount - 1, y: y}].height
			for x := columnCount - 1; x > 0; x-- {
				if forestMap[Coord{x: x, y: y}].height > currentHeight {
					currentHeight = forestMap[Coord{x: x, y: y}].height
					forestMap[Coord{x: x, y: y}].visible = true
				} else {
					continue
				}
			}
		}
	}

	for x := 0; x < columnCount; x++ {
		// Sweep from top
		{
			forestMap[Coord{x: x, y: 0}].visible = true
			currentHeight := forestMap[Coord{x: x, y: 0}].height
			for y := 1; y < rowCount; y++ {
				if forestMap[Coord{x: x, y: y}].height > currentHeight {
					currentHeight = forestMap[Coord{x: x, y: y}].height
					forestMap[Coord{x: x, y: y}].visible = true
				} else {
					continue
				}
			}
		}
		// Sweep from bottom
		{
			forestMap[Coord{x: x, y: rowCount - 1}].visible = true
			currentHeight := forestMap[Coord{x: x, y: rowCount - 1}].height
			for y := rowCount - 1; y > 0; y-- {
				if forestMap[Coord{x: x, y: y}].height > currentHeight {
					currentHeight = forestMap[Coord{x: x, y: y}].height
					forestMap[Coord{x: x, y: y}].visible = true
				} else {
					continue
				}
			}
		}
	}

	ansPart1 := 0
	// println("")
	for y := 0; y < rowCount; y++ {
		for x := 0; x < columnCount; x++ {
			if (forestMap[Coord{x, y}].visible) {
				ansPart1 += 1
				// print(1)
			} else {
				// print(0)
			}
		}
		// println("")
	}
	// println("")
	println("Part 1", ansPart1)

	ansPart2 := 0
	for y := 0; y < rowCount; y++ {
		for x := 0; x < columnCount; x++ {
			if x == 0 || y == 0 || x == columnCount-1 || y == rowCount-1 {
				continue
				// edges will always be 0
			}

			currentHeight := forestMap[Coord{x, y}].height
			scenicScore := 1
			// Sweep to top
			{
				viewingDistance := 0
				for ty := y - 1; ty >= 0; ty-- {
					viewingDistance += 1
					if forestMap[Coord{x: x, y: ty}].height >= currentHeight {
						break
					}
				}
				scenicScore = scenicScore * viewingDistance
			}
			// Sweep to bottom
			{
				viewingDistance := 0
				for ty := y + 1; ty < rowCount; ty++ {
					viewingDistance += 1
					if forestMap[Coord{x: x, y: ty}].height >= currentHeight {
						break
					}
				}
				scenicScore = scenicScore * viewingDistance
			}

			// Sweep to left
			{
				viewingDistance := 0
				for tx := x - 1; tx >= 0; tx-- {
					viewingDistance += 1
					if forestMap[Coord{x: tx, y: y}].height >= currentHeight {
						break
					}
				}
				scenicScore = scenicScore * viewingDistance
			}

			// Sweep to right
			{
				viewingDistance := 0
				for tx := x + 1; tx < columnCount; tx++ {
					viewingDistance += 1
					if forestMap[Coord{x: tx, y: y}].height >= currentHeight {
						break
					}
				}
				scenicScore = scenicScore * viewingDistance
			}

			forestMap[Coord{x, y}].scenicScore = scenicScore

			if scenicScore > ansPart2 {
				ansPart2 = scenicScore
			}
		}
	}

	// println("")
	// for y := 0; y < rowCount; y++ {
	// 	for x := 0; x < columnCount; x++ {
	// 		print(forestMap[Coord{x, y}].scenicScore)
	// 	}
	// 	println("")
	// }
	// println("")
	println("Part 2", ansPart2)

}
