package main

import (
	"strings"
	"utils"
)

/*

Axis positivity in this problem:


              ^ y+
              |
              |
              |
              |
x- <----------0----------> x+
              |
              |
              |
              |
              y-

*/

type Coord struct {
	x int
	y int
}
type Block struct {
	relativeCoords []Coord
	origin         Coord
	height         int
	width          int
}

var rocks = []Block{
	{
		relativeCoords: []Coord{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}, {x: 3, y: 0}},
		origin:         Coord{x: 0, y: 0},
	},
	{
		relativeCoords: []Coord{{x: 1, y: 0}, {x: 0, y: 1}, {x: 1, y: 1}, {x: 2, y: 1}, {x: 1, y: 2}},
		origin:         Coord{x: 0, y: 0},
	}, {
		relativeCoords: []Coord{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}, {x: 2, y: 1}, {x: 2, y: 2}},
		origin:         Coord{x: 0, y: 0},
	},
	{
		relativeCoords: []Coord{{x: 0, y: 0}, {x: 0, y: 1}, {x: 0, y: 2}, {x: 0, y: 3}},
		origin:         Coord{x: 0, y: 0},
	},
	{
		relativeCoords: []Coord{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 1}},
		origin:         Coord{x: 0, y: 0},
	},
}

type CellType int

const (
	CellRock = iota + 1
	CellNone
)

func main() {
	lines := utils.ReadFileToLines("day17.in")
	line := lines[0]
	chars := strings.Split(line, "")
	// Part 1
	// rockCountLimit := 2022
	rockCountLimit := 100_000
	// Part 2
	// rockCountLimit := 1_000_000_000_000
	t := 0
	towerHeight := 0
	areaMap := map[Coord]CellType{}
	movementMap := map[string]Coord{
		"<": {x: -1, y: 0},
		">": {x: 1, y: 0},
	}
	calculateTowerHeight := func() int {
		newTowerHeight := towerHeight
		for y := towerHeight; y < towerHeight+5; y++ {
			isRowHasRock := false
			for x := 0; x <= 6; x++ {
				if areaMap[Coord{x, y}] == CellRock {
					isRowHasRock = true
					break
				}
			}
			if !isRowHasRock {
				newTowerHeight = y
				break
			}
		}
		return newTowerHeight
	}

	printMap := func(yStart int, yEnd int) {
		for y := yEnd; y >= yStart; y-- {
			for x := 0; x <= 6; x++ {
				if areaMap[Coord{x, y}] == CellRock {
					print("#")
				} else {
					print(".")
				}
			}
			println()
		}
	}

	isBlockClear := func(block Block, currentBlock Block) bool {
		currentBlockCoords := map[Coord]bool{}
		for _, relCoord := range currentBlock.relativeCoords {
			coord := absCoord(relCoord, currentBlock.origin)
			currentBlockCoords[coord] = true
		}

		for _, coord := range block.relativeCoords {
			newCoord := absCoord(coord, block.origin)
			if newCoord.x < 0 || newCoord.x > 6 || newCoord.y < 0 {
				// println("not clear (wall)", newCoord.x, newCoord.y)
				return false
			}
			cellType, exist := areaMap[newCoord]
			if !exist {
				// means clear
			}
			if cellType == CellRock && !currentBlockCoords[newCoord] {
				// println("not clear (rock)", newCoord.x, newCoord.y)
				return false
			}
		}
		return true
	}
	clearBlock := func(block Block) {
		for _, coord := range block.relativeCoords {
			newCoord := absCoord(coord, block.origin)
			// println("clear", newCoord.x, newCoord.y)
			areaMap[newCoord] = CellNone
		}
	}
	paintBlock := func(block Block) {
		for _, coord := range block.relativeCoords {
			newCoord := absCoord(coord, block.origin)
			// println("paint", newCoord.x, newCoord.y)
			areaMap[newCoord] = CellRock
		}
	}
	printMap(0, 10)

	for r := 0; r < rockCountLimit; r++ {
		currentRock := rocks[r%len(rocks)]
		currentRock.origin = Coord{
			x: 2,
			y: towerHeight + 3,
		}
		paintBlock(currentRock)
		// println(r, "new", currentRock.origin.x, currentRock.origin.y)
		for isRockAtRest := false; !isRockAtRest; {
			// println(r, "t", t)
			// printMap(0, 9)
			currentChar := chars[t%len(chars)]
			// can it move left / right
			nextBlockLR := currentRock
			nextBlockLR.origin = absCoord(currentRock.origin, movementMap[currentChar])
			if isBlockClear(nextBlockLR, currentRock) {
				// if can, do the movement
				clearBlock(currentRock)
				currentRock = nextBlockLR
				paintBlock(currentRock)
				// println(r, "move", currentChar)
			}

			// printMap(0, 9)
			// can it move down
			nextBlockDown := currentRock
			nextBlockDown.origin = absCoord(currentRock.origin, Coord{x: 0, y: -1})
			if isBlockClear(nextBlockDown, currentRock) {
				clearBlock(currentRock)
				// if can, do it; else break
				currentRock = nextBlockDown
				paintBlock(currentRock)
				// println(r, "move down")
			} else {
				// println(r, "rest")
				isRockAtRest = true
			}
			t++
		}
		newTowerHeight := calculateTowerHeight()
		if r == 1000 {
			println(r, "tower height", towerHeight)
		} else if r >= 1000 {
			towerHeightDiff := newTowerHeight - towerHeight
			print(towerHeightDiff, " ")
		}
		towerHeight = newTowerHeight
		// println(r, "tower height", towerHeight)
	}
	println()
	// println("Ans part 1", towerHeight)
	// printMap(0, towerHeight)
}

func absCoord(relative Coord, origin Coord) Coord {
	return Coord{
		x: relative.x + origin.x,
		y: relative.y + origin.y,
	}
}
