package main

import (
	"strconv"
	"strings"
	"utils"
)

type Coord struct {
	/* distance to the right */
	x int
	/* distance down */
	y int
}
type CellType int

const (
	Air CellType = iota + 1
	Rock
	Sand
)

type Cell struct {
	cellType CellType
}

func main() {

	sandStartCoord := Coord{x: 500, y: 0}
	areaMapPart1 := map[Coord]*Cell{}
	areaMapPart2 := map[Coord]*Cell{}

	maxY := 0

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			areaMapPart1[Coord{x, y}] = &Cell{cellType: Air}
			areaMapPart2[Coord{x, y}] = &Cell{cellType: Air}
		}
	}
	areaMapPart1[sandStartCoord].cellType = Sand
	areaMapPart2[sandStartCoord].cellType = Sand

	lines := utils.ReadFileToLines("day14.in")
	for _, line := range lines {
		if line == "" {
			continue
		}
		coordStrings := strings.Split(line, " -> ")

		var prevCoord Coord

		for k, coordString := range coordStrings {
			coordAxisString := strings.Split(coordString, ",")

			x, _ := strconv.Atoi(coordAxisString[0])
			y, _ := strconv.Atoi(coordAxisString[1])
			areaMapPart1[Coord{x, y}].cellType = Rock
			areaMapPart2[Coord{x, y}].cellType = Rock
			if y > maxY {
				maxY = y
			}

			if k > 0 {
				var lowX = x
				var highX = prevCoord.x
				if prevCoord.x < lowX {
					lowX = prevCoord.x
					highX = x
				}
				var lowY = y
				var highY = prevCoord.y
				if prevCoord.y < lowY {
					lowY = prevCoord.y
					highY = y
				}
				if lowX == highX {
					for ty := lowY; ty <= highY; ty++ {
						areaMapPart1[Coord{x: x, y: ty}].cellType = Rock
						areaMapPart2[Coord{x: x, y: ty}].cellType = Rock
					}
				} else if lowY == highY {
					for tx := lowX; tx <= highX; tx++ {
						areaMapPart1[Coord{x: tx, y: y}].cellType = Rock
						areaMapPart2[Coord{x: tx, y: y}].cellType = Rock
					}
				}
			}

			prevCoord = Coord{x, y}
		}
	}

	for x := 0; x < 1000; x++ {
		areaMapPart2[Coord{x: x, y: maxY + 2}].cellType = Rock
	}

	activeSand := Coord{x: sandStartCoord.x, y: sandStartCoord.y}
	ansPart1 := 0
	ansPart2 := 0

	for t := 0; t < 100_000_000; t++ {
		// println("t", t, "activeSand", activeSand.x, activeSand.y)
		if activeSand.y > maxY+50 {
			// Reached abyss
			break
		}
		{
			// Try 1 step down
			nextCoord := Coord{x: activeSand.x, y: activeSand.y + 1}
			// println("t", t, "nextCoord1", nextCoord.x, nextCoord.y, areaMap[nextCoord].cellType)
			if areaMapPart1[nextCoord].cellType == Air {
				areaMapPart1[activeSand].cellType = Air
				areaMapPart1[nextCoord].cellType = Sand
				activeSand = nextCoord
				continue
			}
		}
		{
			// Try 1 step down & to the left
			nextCoord := Coord{x: activeSand.x - 1, y: activeSand.y + 1}
			// println("t", t, "nextCoord2", nextCoord.x, nextCoord.y, areaMap[nextCoord].cellType)
			if areaMapPart1[nextCoord].cellType == Air {
				areaMapPart1[activeSand].cellType = Air
				areaMapPart1[nextCoord].cellType = Sand
				activeSand = nextCoord
				continue
			}
		}
		{
			// Try 1 step down & to the right
			nextCoord := Coord{x: activeSand.x + 1, y: activeSand.y + 1}
			// println("t", t, "nextCoord3", nextCoord.x, nextCoord.y, areaMap[nextCoord].cellType)
			if areaMapPart1[nextCoord].cellType == Air {
				areaMapPart1[activeSand].cellType = Air
				areaMapPart1[nextCoord].cellType = Sand
				activeSand = nextCoord
				continue
			}
		}
		// Comes to rest
		// Reset activeSand
		// println("Rest!", activeSand.x, activeSand.y)
		activeSand = Coord{x: sandStartCoord.x, y: sandStartCoord.y}
		ansPart1 += 1
	}

	println("ansPart1", ansPart1)

	activeSand = Coord{x: sandStartCoord.x, y: sandStartCoord.y}

	for t := 0; t < 1000_000_000; t++ {
		// println("t", t, "activeSand", activeSand.x, activeSand.y)

		{
			// Try 1 step down
			nextCoord := Coord{x: activeSand.x, y: activeSand.y + 1}
			// println("t", t, "nextCoord1", nextCoord.x, nextCoord.y, areaMapPart2[nextCoord].cellType)
			if areaMapPart2[nextCoord].cellType == Air {
				areaMapPart2[activeSand].cellType = Air
				areaMapPart2[nextCoord].cellType = Sand
				activeSand = nextCoord
				continue
			}
		}
		{
			// Try 1 step down & to the left
			nextCoord := Coord{x: activeSand.x - 1, y: activeSand.y + 1}
			// println("t", t, "nextCoord2", nextCoord.x, nextCoord.y, areaMapPart2[nextCoord].cellType)
			if areaMapPart2[nextCoord].cellType == Air {
				areaMapPart2[activeSand].cellType = Air
				areaMapPart2[nextCoord].cellType = Sand
				activeSand = nextCoord
				continue
			}
		}
		{
			// Try 1 step down & to the right
			nextCoord := Coord{x: activeSand.x + 1, y: activeSand.y + 1}
			// println("t", t, "nextCoord3", nextCoord.x, nextCoord.y, areaMapPart2[nextCoord].cellType)
			if areaMapPart2[nextCoord].cellType == Air {
				areaMapPart2[activeSand].cellType = Air
				areaMapPart2[nextCoord].cellType = Sand
				activeSand = nextCoord
				continue
			}
		}

		if activeSand.x == sandStartCoord.x && activeSand.y == sandStartCoord.y {
			// Reached start coord
			ansPart2 += 1
			break
		}

		// Comes to rest
		// Reset activeSand
		// println("Rest!", activeSand.x, activeSand.y)
		activeSand = Coord{x: sandStartCoord.x, y: sandStartCoord.y}
		ansPart2 += 1
	}

	println("ansPart2", ansPart2)
}
