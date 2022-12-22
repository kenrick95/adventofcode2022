package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

type CellType int
type Direction int
type TurnType int

const (
	Offside CellType = iota
	OpenTile
	SolidWall
)
const (
	DirRight Direction = 0
	DirDown  Direction = 1
	DirLeft  Direction = 2
	DirUp    Direction = 3
)
const (
	TurnLeft  TurnType = -1
	TurnRight TurnType = 1
)

type Cell struct {
	content CellType
}
type Coord struct {
	x int
	y int
}
type State struct {
	coord     Coord
	direction Direction
}
type Command struct {
	amount int
	turn   TurnType
}

func main() {
	areaMap := map[Coord]Cell{}
	lines := utils.ReadFileToLinesNoTrim("day22.real.in")
	startingCoord := Coord{x: 0, y: 0}
	hasFoundStartingCoord := false
	hasReachedCommandsSection := false
	re := regexp.MustCompile(`(?m)((\d+)(L|R))`)
	commands := []Command{}

	xMin := 0
	yMin := 0
	xMax := 0
	yMax := 0

	for y, line := range lines {
		if strings.TrimSpace(line) == "" {
			hasReachedCommandsSection = true
			continue
		}

		if hasReachedCommandsSection {
			for _, match := range re.FindAllStringSubmatch(line, -1) {
				amount, _ := strconv.Atoi(match[2])
				turnString := match[3]
				turn := TurnLeft
				if turnString == "L" {
					turn = TurnLeft
				} else if turnString == "R" {
					turn = TurnRight
				}
				commands = append(commands, Command{
					amount: amount,
					turn:   turn,
				})
			}

		} else {

			chars := strings.Split(line, "")
			for x, char := range chars {
				coord := Coord{x, y}
				cellType := Offside
				if char == "." {
					cellType = OpenTile
					if y == 0 && !hasFoundStartingCoord {
						hasFoundStartingCoord = true
						startingCoord = coord
					}
				} else if char == "#" {
					cellType = SolidWall
				} else if char != " " {
					continue
				}

				xMax = utils.Max(xMax, coord.x)
				yMax = utils.Max(yMax, coord.y)
				areaMap[coord] = Cell{
					content: cellType,
				}
			}
		}
	}

	/* At column "x", what is the lowest "y" */
	xBoundaryLo := map[int]int{}
	/* At column "x", what is the highest "y" */
	xBoundaryHi := map[int]int{}
	/* At row "y", what is the lowest "x" */
	yBoundaryLo := map[int]int{}
	/* At row "y", what is the highest "x" */
	yBoundaryHi := map[int]int{}

	for x := xMin - 1; x <= xMax+1; x++ {
		xBoundaryLo[x] = yMax
		xBoundaryHi[x] = yMin
		for y := yMin; y <= yMax; y++ {
			_, exist := areaMap[Coord{x, y}]
			if !exist {
				areaMap[Coord{x, y}] = Cell{content: Offside}
			}
			cellType := areaMap[Coord{x, y}].content
			if cellType == OpenTile || cellType == SolidWall {
				xBoundaryLo[x] = utils.Min(xBoundaryLo[x], y)
				xBoundaryHi[x] = utils.Max(xBoundaryHi[x], y)
			}
		}
	}

	for y := yMin - 1; y <= yMax+1; y++ {
		yBoundaryLo[y] = xMax
		yBoundaryHi[y] = xMin
		for x := xMin; x <= xMax; x++ {
			_, exist := areaMap[Coord{x, y}]
			if !exist {
				areaMap[Coord{x, y}] = Cell{content: Offside}
			}
			cellType := areaMap[Coord{x, y}].content
			if cellType == OpenTile || cellType == SolidWall {
				yBoundaryLo[y] = utils.Min(yBoundaryLo[y], x)
				yBoundaryHi[y] = utils.Max(yBoundaryHi[y], x)
			}
		}
	}

	fmt.Printf("commands: %v\n", commands)
	fmt.Printf("startingCoord: %v\n", startingCoord)

	deltas := map[Direction]Coord{
		DirDown:  {x: 0, y: 1},
		DirUp:    {x: 0, y: -1},
		DirLeft:  {x: -1, y: 0},
		DirRight: {x: 1, y: 0},
	}
	getNextDirection := func(currentDirection Direction, turn TurnType) Direction {
		if turn == TurnLeft {
			switch currentDirection {
			case DirDown:
				return DirRight
			case DirUp:
				return DirLeft
			case DirLeft:
				return DirDown
			case DirRight:
				return DirUp
			}
		} else if turn == TurnRight {
			switch currentDirection {
			case DirDown:
				return DirLeft
			case DirUp:
				return DirRight
			case DirLeft:
				return DirUp
			case DirRight:
				return DirDown
			}
		}
		return DirDown
	}

	walk := func(initialState State, isPart2 bool) int {
		currentState := initialState
		for _, cmd := range commands {

			// fmt.Printf("cmd: %v %v\n", i, cmd)
			// fmt.Printf("currentState: %v\n", currentState)

			for count := 0; count < cmd.amount; count++ {
				currentDelta := deltas[currentState.direction]
				nextCoord := Coord{
					x: currentState.coord.x + currentDelta.x,
					y: currentState.coord.y + currentDelta.y,
				}
				nextDirection := currentState.direction
				// fmt.Printf("nextCoord: %v\n", nextCoord)

				nextCell, _ := areaMap[nextCoord]

				// fmt.Printf("nextCell: %v\n", nextCell)

				if nextCell.content == Offside {
					// wrap~
					if !isPart2 {
						// Part 1 logic
						if currentDelta.x == 0 {
							// wrap y
							if currentState.coord.y == xBoundaryLo[currentState.coord.x] {
								nextCoord.y = xBoundaryHi[currentState.coord.x]
							} else if currentState.coord.y == xBoundaryHi[currentState.coord.x] {
								nextCoord.y = xBoundaryLo[currentState.coord.x]
							}
						} else if currentDelta.y == 0 {
							// wrap x
							if currentState.coord.x == yBoundaryLo[currentState.coord.y] {
								nextCoord.x = yBoundaryHi[currentState.coord.y]
							} else if currentState.coord.x == yBoundaryHi[currentState.coord.y] {
								nextCoord.x = yBoundaryLo[currentState.coord.y]
							}
						}
					} else {
						// Part 2 logic
						// TODO: Hmmm quite hard
					}

					// fmt.Printf("nextCoord (wrapped): %v\n", nextCoord)
					nextCell, _ = areaMap[nextCoord]
					// fmt.Printf("nextCell (wrapped): %v\n", nextCell)
				}

				if nextCell.content == SolidWall {
					break
				} else {
					currentState.coord = nextCoord
					currentState.direction = nextDirection
				}
			}

			nextDirection := getNextDirection(currentState.direction, cmd.turn)

			currentState.direction = nextDirection
		}

		fmt.Printf("currentState: %v\n", currentState)

		ans := 1000*(currentState.coord.y+1) + 4*(currentState.coord.x+1) + int(currentState.direction)
		return ans
	}

	// Part 1
	fmt.Printf("ansPart1: %v\n", walk(State{
		coord:     startingCoord,
		direction: DirRight,
	}))
}
