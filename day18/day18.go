package main

import (
	"strconv"
	"strings"
	"utils"
)

type Coord struct {
	x int
	y int
	z int
}
type CellType int

const (
	CellEmpty CellType = iota + 1
	CellRock
	CellEmptyInside
	CellEmptyOutside
)

type BfsNode struct {
	coord Coord
}

func main() {
	scanMap := map[Coord]CellType{}
	deltas := []Coord{
		{0, 0, 1},
		{0, 0, -1},
		{0, 1, 0},
		{0, -1, 0},
		{1, 0, 0},
		{-1, 0, 0},
	}
	lines := utils.ReadFileToLines("day18.in")
	// For real input
	// max 19 19 19
	// min 0 0 0
	maxX, maxY, maxZ := 0, 0, 0
	minX, minY, minZ := 100_000, 100_000, 100_000
	for _, line := range lines {
		if line == "" {
			continue
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		scanMap[Coord{x, y, z}] = CellRock
		maxX, maxY, maxZ = utils.Max(x, maxX), utils.Max(y, maxY), utils.Max(z, maxZ)
		minX, minY, minZ = utils.Min(x, minX), utils.Min(y, minY), utils.Min(z, minZ)
	}
	minX -= 1
	minY -= 1
	minZ -= 1
	maxX += 1
	maxY += 1
	maxZ += 1
	println("max", maxX, maxY, maxZ)
	println("min", minX, minY, minZ)

	for y := minY; y <= maxY; y++ {
		for z := minZ; z <= maxZ; z++ {
			for x := minX; x <= maxX; x++ {
				checkCoord := Coord{x, y, z}
				_, coordExist := scanMap[checkCoord]
				if !coordExist {
					scanMap[checkCoord] = CellEmpty
				}
			}
		}
	}

	ansPart1 := 0
	for coord := range scanMap {
		if scanMap[coord] != CellRock {
			continue
		}
		for _, delta := range deltas {
			checkCoord := Coord{
				x: coord.x + delta.x,
				y: coord.y + delta.y,
				z: coord.z + delta.z,
			}
			cellType, _ := scanMap[checkCoord]
			if cellType == CellEmpty {
				ansPart1 += 1
			}
		}
	}
	println("ansPart1", ansPart1)

	bfsQueue := []BfsNode{}
	bfsVisitedMap := map[Coord]bool{}

	initCoords := []Coord{
		{minX, minY, minZ},
		{minX, minY, maxZ},
		{minX, maxY, minZ},
		{minX, maxY, maxZ},
		{maxX, minY, minZ},
		{maxX, minY, maxZ},
		{maxX, maxY, minZ},
		{maxX, maxY, maxZ},
	}
	for _, checkCoord := range initCoords {
		bfsQueue = append(bfsQueue, BfsNode{coord: checkCoord})
		bfsVisitedMap[checkCoord] = true
	}

	// meow := 0
	for len(bfsQueue) > 0 {
		var node BfsNode
		node, bfsQueue = bfsQueue[0], bfsQueue[1:]

		for _, delta := range deltas {
			nextCoord := Coord{
				x: node.coord.x + delta.x,
				y: node.coord.y + delta.y,
				z: node.coord.z + delta.z,
			}
			if nextCoord.x >= minX &&
				nextCoord.y >= minY &&
				nextCoord.z >= minZ &&
				nextCoord.x <= maxX &&
				nextCoord.y <= maxY &&
				nextCoord.z <= maxZ && !bfsVisitedMap[nextCoord] {
				cellType, _ := scanMap[nextCoord]
				if cellType == CellEmpty {
					// println("marking as outside", nextCoord.x, nextCoord.y, nextCoord.z)
					scanMap[nextCoord] = CellEmptyOutside
					bfsQueue = append(bfsQueue, BfsNode{coord: nextCoord})
					bfsVisitedMap[nextCoord] = true
					// } else if cellType == CellRock {
					// 	meow += 1
				}
			}
		}
	}
	// println("meow", meow)

	for y := minY; y <= maxY; y++ {
		for z := minZ; z <= maxZ; z++ {
			for x := minX; x <= maxX; x++ {
				checkCoord := Coord{x, y, z}
				cellType := scanMap[checkCoord]
				if cellType == CellEmpty {
					scanMap[checkCoord] = CellEmptyInside
					// println("marking as inside", checkCoord.x, checkCoord.y, checkCoord.z)
				}
				// print(scanMap[Coord{x, z, y}], " ")
			}
			// print("; ")
		}
		// println()
	}

	ansPart2 := 0
	for coord := range scanMap {
		if scanMap[coord] != CellRock {
			continue
		}
		// println("coord", coord.x, coord.y, coord.z)
		for _, delta := range deltas {
			checkCoord := Coord{
				x: coord.x + delta.x,
				y: coord.y + delta.y,
				z: coord.z + delta.z,
			}
			cellType, _ := scanMap[checkCoord]
			if cellType == CellEmptyOutside {
				// println("coord2", checkCoord.x, checkCoord.y, checkCoord.z, coordExist, cellType)
				ansPart2 += 1
			}
		}
	}
	println("ansPart2", ansPart2)
}
