package main

import (
	"strconv"
	"strings"
	"utils"
)

/*

1. Touching: H & T must be adjacent (hor, ver, diag) or even overlapping
2. Same Row/Column: if H & T is two steps away (hor, ver), T must also take 1 step in that direction (hor or ver) to be touching to H
3. Else: (i.e. H & T are not touching & not in same row/column), T must take 1 step diag to be touching to H
4. H & T starts at same position, overallping

*/

type Coord struct {
	x int
	y int
}

func main() {
	dirMap := map[string]Coord{
		"D": {x: 0, y: 1},
		"U": {x: 0, y: -1},
		"L": {x: -1, y: 0},
		"R": {x: 1, y: 0},
	}
	lines := utils.ReadFileToLines("day09.in")

	// Part 1
	{
		posHead := Coord{x: 0, y: 0}
		posTail := Coord{x: 0, y: 0}
		tailVisitedMap := map[Coord]bool{}
		tailVisitedMap[posTail] = true
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.Split(line, " ")
			dir := parts[0]
			delta, _ := strconv.Atoi(parts[1])
			deltaCoord := dirMap[dir]

			// println("> ", line)

			for i := 0; i < delta; i++ {
				// Update head
				posHead = Coord{
					x: posHead.x + deltaCoord.x,
					y: posHead.y + deltaCoord.y,
				}
				// println("Head:", posHead.x, posHead.y)

				// Update tail
				dx := posHead.x - posTail.x
				dy := posHead.y - posTail.y
				dxAbs := abs(dx)
				dyAbs := abs(dy)
				dxUnit := unit(dx)
				dyUnit := unit(dy)
				if dxAbs <= 1 && dyAbs <= 1 {
					// 1. is touching? --> nothing
					continue
				} else if dxAbs == 0 || dyAbs == 0 {
					// 2. is same row/col --> move hor/col

					posTail = Coord{
						x: posTail.x + dxUnit,
						y: posTail.y + dyUnit,
					}

				} else {
					// 3. other case --> move diag
					posTail = Coord{
						x: posTail.x + dxUnit,
						y: posTail.y + dyUnit,
					}
				}
				// println("Tail:", posTail.x, posTail.y)

				tailVisitedMap[posTail] = true
			}

		}

		ansPart1 := len(tailVisitedMap)
		// println("Visited:")
		// for coord := range tailVisitedMap {
		// 	println(coord.x, coord.y)
		// }

		println("Part 1:", ansPart1)

	}

	// Part 2: 10 knots
	{
		var posKnots []Coord
		for i := 0; i < 10; i++ {
			posKnots = append(posKnots, Coord{x: 0, y: 0})
		}
		tailVisitedMap := map[Coord]bool{}
		tailVisitedMap[Coord{x: 0, y: 0}] = true

		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.Split(line, " ")
			dir := parts[0]
			delta, _ := strconv.Atoi(parts[1])
			deltaCoord := dirMap[dir]

			// println("> ", line)

			for step := 0; step < delta; step++ {
				// Update head
				posKnots[0] = Coord{
					x: posKnots[0].x + deltaCoord.x,
					y: posKnots[0].y + deltaCoord.y,
				}
				// println("Head:", posKnots[0].x, posKnots[0].y)

				// Update other non-head knots
				for i := 1; i < 10; i++ {
					dx := posKnots[i-1].x - posKnots[i].x
					dy := posKnots[i-1].y - posKnots[i].y
					dxAbs := abs(dx)
					dyAbs := abs(dy)
					dxUnit := unit(dx)
					dyUnit := unit(dy)
					if dxAbs <= 1 && dyAbs <= 1 {
						// 1. is touching? --> nothing
						continue
					} else if dxAbs == 0 || dyAbs == 0 {
						// 2. is same row/col --> move hor/col

						posKnots[i] = Coord{
							x: posKnots[i].x + dxUnit,
							y: posKnots[i].y + dyUnit,
						}

					} else {
						// 3. other case --> move diag
						posKnots[i] = Coord{
							x: posKnots[i].x + dxUnit,
							y: posKnots[i].y + dyUnit,
						}
					}
					// println("Knot", i, " pos:", posKnots[i].x, posKnots[i].y)
				}

				tailVisitedMap[posKnots[9]] = true
			}

		}

		ansPart2 := len(tailVisitedMap)
		// println("Visited:")
		// for coord := range tailVisitedMap {
		// 	println(coord.x, coord.y)
		// }

		println("Part 2:", ansPart2)
	}

}
func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func unit(val int) int {
	if val == 0 {
		return 0
	} else if val < 0 {
		return -1
	}
	return 1
}
