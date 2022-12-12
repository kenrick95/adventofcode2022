package main

import (
	"strings"

	"utils"
)

type Coord struct {
	x int
	y int
}
type Node struct {
	coord Coord
	steps int
}

func main() {
	lines := utils.ReadFileToLines("day12.in")
	heightMap := map[Coord]int{}
	// Part 1: start at "S"
	// Part 2: start at all "a" ("S" is considered "a")
	var part1CoordStart Coord
	var part2CoordsStart []Coord
	var coordEnd Coord
	for y, line := range lines {
		if line == "" {
			continue
		}
		chars := strings.Split(line, "")
		for x, char := range chars {
			heightChar := char
			if char == "S" {
				part1CoordStart = Coord{
					x: x,
					y: y,
				}
				part2CoordsStart = append(part2CoordsStart, Coord{
					x: x,
					y: y,
				})
				heightChar = "a"
			} else if char == "E" {
				coordEnd = Coord{
					x: x,
					y: y,
				}
				heightChar = "z"
			} else if char == "a" {
				part2CoordsStart = append(part2CoordsStart, Coord{
					x: x,
					y: y,
				})
			}
			heightMap[Coord{x: x, y: y}] = int([]rune(heightChar)[0] - []rune("a")[0])
		}
	}

	directions := []Coord{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: -1, y: 0},
		{x: 1, y: 0},
	}

	bfs := func(startCoords []Coord) int {
		var stepsToEnd int
		queue := []Node{}
		for _, coordStart := range startCoords {
			queue = append(queue, Node{
				coord: coordStart,
				steps: 0,
			})
		}

		visitedMap := map[Coord]bool{}
		for coord := range heightMap {
			visitedMap[coord] = false
		}

		{
			for len(queue) > 0 {
				var node Node
				node, queue = queue[0], queue[1:]

				// get next nodes
				nextCoords := []Coord{}
				for _, dir := range directions {
					nextCoord := Coord{
						x: node.coord.x + dir.x,
						y: node.coord.y + dir.y,
					}
					_, exist := heightMap[nextCoord]
					if !exist {
						continue
					}

					if heightMap[nextCoord] > heightMap[node.coord]+1 {
						continue
					}

					nextCoords = append(nextCoords, nextCoord)
				}
				for _, nextCoord := range nextCoords {
					if !visitedMap[nextCoord] {
						visitedMap[nextCoord] = true

						if nextCoord.x == coordEnd.x && nextCoord.y == coordEnd.y {
							// reached
							stepsToEnd = node.steps + 1
							break
						}

						queue = append(queue, Node{
							coord: nextCoord,
							steps: node.steps + 1,
						})
					}
				}

			}
		}

		return stepsToEnd
	}

	println("Part 1", bfs([]Coord{part1CoordStart}))
	println("Part 2", bfs(part2CoordsStart))

}
