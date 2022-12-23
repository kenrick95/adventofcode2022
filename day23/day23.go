package main

import (
	"fmt"
	"strings"
	"utils"
)

type Coord struct {
	x int
	y int
}
type CellType int

const (
	Empty CellType = iota + 1
	ElfOccupied
)

type ElfNode struct {
	id        int
	coord     Coord
	nextCoord Coord
}

func main() {
	areaMap := map[Coord]CellType{}
	initialElfMap := map[Coord]bool{}
	elfNodes := map[int]*ElfNode{}
	elfCount := 0

	lines := utils.ReadFileToLines("day23.real.in")
	deltas := map[string]Coord{
		"N":  {x: 0, y: -1},
		"S":  {x: 0, y: 1},
		"E":  {x: 1, y: 0},
		"W":  {x: -1, y: 0},
		"NE": {x: 1, y: -1},
		"NW": {x: -1, y: -1},
		"SE": {x: 1, y: 1},
		"SW": {x: -1, y: 1},
	}

	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, ch := range strings.Split(line, "") {
			cellType := Empty
			coord := Coord{x, y}
			if ch == "." {
				cellType = Empty
			} else if ch == "#" {
				cellType = ElfOccupied
				initialElfMap[coord] = true
				elfNodes[elfCount] = &ElfNode{
					id:        elfCount,
					coord:     coord,
					nextCoord: coord,
				}
				elfCount += 1
			}
			areaMap[coord] = cellType
		}
	}

	copyMap := func(initialElfMap map[Coord]bool) map[Coord]bool {
		elfMap := map[Coord]bool{}
		for key, value := range initialElfMap {
			elfMap[key] = value
		}
		return elfMap
	}

	elfMap := copyMap(initialElfMap)

	getMinMaxCoords := func() (int, int, int, int) {

		minX := elfNodes[0].coord.x
		maxX := elfNodes[0].coord.x
		minY := elfNodes[0].coord.y
		maxY := elfNodes[0].coord.y

		for i := 1; i < elfCount; i++ {
			coord := elfNodes[i].coord

			minX = utils.Min(minX, coord.x)
			maxX = utils.Max(maxX, coord.x)
			minY = utils.Min(minY, coord.y)
			maxY = utils.Max(maxY, coord.y)
		}
		return minX, maxX, minY, maxY
	}

	debug := func() {
		fmt.Println()
		fmt.Println()
		minX, maxX, minY, maxY := getMinMaxCoords()
		fmt.Println("x: ", minX, "  -  ", maxX, "; y: ", minY, "  -  ", maxY)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if elfMap[Coord{x, y}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()

	}

	getNextCoord := func(coord Coord, direction string) Coord {
		delta := deltas[direction]
		return Coord{
			x: coord.x + delta.x,
			y: coord.y + delta.y,
		}
	}

	round := func(roundCount int) bool {
		nextCoordElfIds := map[Coord][]int{}
		isElfMoving := false

		getNextDirection := func(coord Coord, roundCount int) string {
			hasElfInDirection := map[string]bool{}
			hasElfInSurroundings := false
			for direction, _ := range deltas {
				hasElfInDir := elfMap[getNextCoord(coord, direction)]
				hasElfInDirection[direction] = hasElfInDir
				if hasElfInDir {
					hasElfInSurroundings = true
				}
			}

			if !hasElfInSurroundings {
				return ""
			}

			ruleResults := map[string]bool{
				"N": !hasElfInDirection["N"] && !hasElfInDirection["NE"] && !hasElfInDirection["NW"],
				"S": !hasElfInDirection["S"] && !hasElfInDirection["SE"] && !hasElfInDirection["SW"],
				"W": !hasElfInDirection["W"] && !hasElfInDirection["NW"] && !hasElfInDirection["SW"],
				"E": !hasElfInDirection["E"] && !hasElfInDirection["NE"] && !hasElfInDirection["SE"],
			}
			rules := []string{
				"N",
				"S",
				"W",
				"E",
			}

			ruleStart := roundCount % 4
			for r := ruleStart; r < 4; r++ {
				if ruleResults[rules[r]] {
					return rules[r]
				}
			}
			for r := 0; r < ruleStart; r++ {
				if ruleResults[rules[r]] {
					return rules[r]
				}
			}

			return ""
		}

		// fmt.Println("\nRound", roundCount, "pre-move:")
		// debug()

		// 1: Propose next coord
		for i := 0; i < elfCount; i++ {
			elfNode := elfNodes[i]

			elfNode.nextCoord = elfNode.coord
			direction := getNextDirection(elfNode.coord, roundCount)
			// fmt.Println("Round", roundCount, elfNode.id, elfNode.coord, direction)
			if direction == "" {
				continue
			} else {
				nextCoord := getNextCoord(elfNode.coord, direction)
				elfNode.nextCoord = nextCoord

				_, exist := nextCoordElfIds[nextCoord]
				if !exist {
					nextCoordElfIds[nextCoord] = []int{}
				}

				nextCoordElfIds[nextCoord] = append(nextCoordElfIds[nextCoord], elfNode.id)
			}

		}

		// 2: Move~~
		for nextCoord, elfIds := range nextCoordElfIds {
			if len(elfIds) == 1 {
				isElfMoving = true
				elfId := elfIds[0]
				elfNode := elfNodes[elfId]
				elfMap[elfNode.coord] = false
				elfMap[nextCoord] = true
				elfNode.coord = nextCoord
			} else {
				// none moves
			}
		}

		// fmt.Println("Round", roundCount, "post-move:")
		// debug()
		// fmt.Println("Round", roundCount, "post-move elf positions:")
		// for i := 0; i < len(elfNodes); i++ {
		// 	e := elfNodes[i]
		// 	fmt.Printf("%v: %v; ", e.id, e.coord)
		// }
		return isElfMoving
	}

	{
		// for i := 0; i < len(elfNodes); i++ {
		// 	e := elfNodes[i]
		// 	fmt.Printf("%v: %v; ", e.id, e.coord)
		// }
		debug()
		for r := 0; r < 10_000; r++ {

			// fmt.Print("\nSTART Round ", r, " \n\n")

			isElfMoving := round(r)

			// fmt.Print("\nEND r ", r, " \n\n")
			if !isElfMoving {
				fmt.Println("ansPart2", r+1)
				break
			}
			if r == 9 {
				minX, maxX, minY, maxY := getMinMaxCoords()

				area := utils.Abs(maxX-minX+1) * utils.Abs(maxY-minY+1)
				ansPart1 := area - len(elfNodes)

				fmt.Println("ansPart1", ansPart1)
			}
		}
	}
}
