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

	quadrants := map[Coord]int{}

	{
		qCount := 50
		// Not proud of this hardcoding
		for y := 0; y < qCount; y++ {
			for x := 50; x < 50+qCount; x++ {
				quadrants[Coord{x, y}] = 1
			}
		}
		for y := 0; y < qCount; y++ {
			for x := 100; x < 100+qCount; x++ {
				quadrants[Coord{x, y}] = 2
			}
		}
		for y := 50; y < 50+qCount; y++ {
			for x := 50; x < 50+qCount; x++ {
				quadrants[Coord{x, y}] = 3
			}
		}

		for y := 100; y < 100+qCount; y++ {
			for x := 0; x < 0+qCount; x++ {
				quadrants[Coord{x, y}] = 5
			}
		}
		for y := 100; y < 100+qCount; y++ {
			for x := 50; x < 50+qCount; x++ {
				quadrants[Coord{x, y}] = 6
			}
		}
		for y := 150; y < 150+qCount; y++ {
			for x := 0; x < 0+qCount; x++ {
				quadrants[Coord{x, y}] = 4
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

			fmt.Printf("cmd: %v\n", cmd)

			for count := 0; count < cmd.amount; count++ {
				fmt.Printf("currentState: %v\n", currentState)
				currentDelta := deltas[currentState.direction]
				nextCoord := Coord{
					x: currentState.coord.x + currentDelta.x,
					y: currentState.coord.y + currentDelta.y,
				}
				nextDirection := currentState.direction
				fmt.Printf("nextCoord: %v\n", nextCoord)

				nextCell, _ := areaMap[nextCoord]

				fmt.Printf("nextCell: %v\n", nextCell)

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
						// TODO: Hmmm extremely dirty and complicated and bug-prone (dead)
						currentQ := quadrants[currentState.coord]
						nextQ := 0
						fmt.Printf("wrapping currentQ: %v\n", currentQ)

						switch currentQ {
						case 1:
							{
								if currentState.direction == DirDown {
									nextQ = 3
								} else if currentState.direction == DirUp {
									nextQ = 4
									nextDirection = DirRight
									nextCoord = Coord{
										x: 0,
										y: currentState.coord.x - 50 + 150,
									}
								} else if currentState.direction == DirRight {
									nextQ = 2
								} else if currentState.direction == DirLeft {
									nextQ = 5
									nextDirection = DirRight
									nextCoord = Coord{
										x: 0,
										y: -currentState.coord.y + 149,
									}
								}
							}
						case 2:
							{
								if currentState.direction == DirDown {
									nextQ = 3
									nextDirection = DirLeft
									nextCoord = Coord{
										x: 99,
										y: currentState.coord.x - 100 + 50,
									}
								} else if currentState.direction == DirUp {
									nextQ = 4
									nextDirection = DirUp
									nextCoord = Coord{
										x: currentState.coord.x - 100,
										y: 199,
									}
								} else if currentState.direction == DirRight {
									nextQ = 6
									nextDirection = DirLeft
									nextCoord = Coord{
										x: 99,
										y: -currentState.coord.y + 149,
									}
								} else if currentState.direction == DirLeft {
									nextQ = 1
								}
							}
						case 3:
							{
								if currentState.direction == DirDown {
									nextQ = 6
								} else if currentState.direction == DirUp {
									nextQ = 1
								} else if currentState.direction == DirRight {
									nextQ = 2
									nextDirection = DirUp
									nextCoord = Coord{
										x: currentState.coord.y - 50 + 100,
										y: 49,
									}
								} else if currentState.direction == DirLeft {
									nextQ = 5
									nextDirection = DirDown
									nextCoord = Coord{
										x: currentState.coord.y - 50,
										y: 100,
									}
								}
							}

						case 4:
							{
								if currentState.direction == DirDown {
									nextQ = 2
									nextDirection = DirDown
									nextCoord = Coord{
										x: currentState.coord.x + 100,
										y: 0,
									}
								} else if currentState.direction == DirUp {
									nextQ = 5
								} else if currentState.direction == DirRight {
									nextQ = 6
									nextDirection = DirUp
									nextCoord = Coord{
										x: currentState.coord.y - 150 + 50,
										y: 149,
									}
								} else if currentState.direction == DirLeft {
									nextQ = 1
									nextDirection = DirDown
									nextCoord = Coord{
										x: currentState.coord.y - 150 + 50,
										y: 0,
									}
								}
							}

						case 5:
							{
								if currentState.direction == DirDown {
									nextQ = 4
								} else if currentState.direction == DirUp {
									nextQ = 3
									nextDirection = DirRight
									nextCoord = Coord{
										x: 50,
										y: currentState.coord.x + 50,
									}
								} else if currentState.direction == DirRight {
									nextQ = 6
								} else if currentState.direction == DirLeft {
									nextQ = 1
									nextDirection = DirRight
									nextCoord = Coord{
										x: 50,
										y: -(currentState.coord.y - 100) + 49,
									}
								}
							}

						case 6:
							{
								if currentState.direction == DirDown {
									nextQ = 4
									nextDirection = DirLeft
									nextCoord = Coord{
										x: 49,
										y: currentState.coord.x - 50 + 150,
									}
								} else if currentState.direction == DirUp {
									nextQ = 3
								} else if currentState.direction == DirRight {
									nextQ = 2
									nextDirection = DirLeft
									nextCoord = Coord{
										x: 149,
										y: -(currentState.coord.y - 100) + 49,
									}
								} else if currentState.direction == DirLeft {
									nextQ = 5
								}
							}
						}

						fmt.Printf("nextQ: %v\n", nextQ)

					}

					fmt.Printf("nextCoord (wrapped): %v\n", nextCoord)
					fmt.Printf("nextDirection (wrapped): %v\n", nextDirection)
					nextCell, _ = areaMap[nextCoord]
					fmt.Printf("nextCell (wrapped): %v\n", nextCell)
				}

				if nextCell.content == SolidWall {
					fmt.Printf("Hit wall, breaking\n")
					break
				} else {
					fmt.Printf("OK\n")
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
	// fmt.Printf("ansPart1: %v\n", walk(State{
	// 	coord:     startingCoord,
	// 	direction: DirRight,
	// }, false))

	// Part 2
	// That's not the right answer; your answer is too high.  (You guessed 169153.)
	// That's not the right answer; your answer is too high.  (You guessed 124314.)
	fmt.Printf("ansPart2: %v\n", walk(State{
		coord:     startingCoord,
		direction: DirRight,
	}, true))
}
