package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Coord struct {
	x int
	y int
}

const (
	DirUp    = 1
	DirDown  = 2
	DirLeft  = 3
	DirRight = 4
)
const (
	CellWall = iota + 1
	CellEmpty
	CellBlizzard
)

type Blizzard struct {
	id        int
	coord     Coord
	direction int
}

type State struct {
	coord          Coord
	blizzardStates map[int]Blizzard
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    State // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value State, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	lines := utils.ReadFileToLines("day24.in")
	areaMap := map[Coord]int{}
	blizzardMap := map[int]Blizzard{}
	blizzardCount := 0
	coordStart := Coord{x: 0, y: 0}
	coordEnd := Coord{x: 0, y: 0}
	deltas := map[int]Coord{
		DirUp:    {x: 0, y: -1},
		DirDown:  {x: 0, y: 1},
		DirRight: {x: 1, y: 0},
		DirLeft:  {x: -1, y: 0},
	}
	minNonWallXForY := map[int]int{}
	maxNonWallXForY := map[int]int{}
	minNonWallYForX := map[int]int{}
	maxNonWallYForX := map[int]int{}
	minX, maxX, minY, maxY := 0, 0, 0, 0

	copyMap := func(initialMap map[int]Blizzard) map[int]Blizzard {
		bMap := map[int]Blizzard{}
		for key, value := range initialMap {
			bMap[key] = value
		}
		return bMap
	}
	getNextCoord := func(coord Coord, dir int) Coord {
		delta := deltas[dir]
		return Coord{
			x: coord.x + delta.x,
			y: coord.y + delta.y,
		}
	}
	getNextBlizzardStates := func(initialMap map[int]Blizzard) map[int]Blizzard {
		blizzardMap := copyMap(initialMap)
		for _, blizzard := range blizzardMap {
			nextCoord := getNextCoord(blizzard.coord, blizzard.direction)

			if areaMap[nextCoord] == CellWall {
				switch blizzard.direction {
				case DirLeft:
					{
						nextCoord.x = maxNonWallXForY[nextCoord.y]
					}
				case DirRight:
					{
						nextCoord.x = minNonWallXForY[nextCoord.y]
					}
				case DirUp:
					{
						nextCoord.y = maxNonWallYForX[nextCoord.x]
					}
				case DirDown:
					{
						nextCoord.y = minNonWallYForX[nextCoord.x]
					}
				}

			}

			blizzard.coord = nextCoord
		}

		return blizzardMap
	}

	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, ch := range strings.Split(line, "") {
			coord := Coord{x, y}
			cellType := CellEmpty
			blizzard := Blizzard{coord: coord, direction: 0}
			switch ch {
			case "#":
				{
					cellType = CellWall
				}
			case ".":
				{
					cellType = CellEmpty
				}
			case ">":
				{
					cellType = CellBlizzard
					blizzard.direction = DirRight
				}
			case "<":
				{
					cellType = CellBlizzard
					blizzard.direction = DirLeft
				}
			case "^":
				{
					cellType = CellBlizzard
					blizzard.direction = DirUp
				}
			case "v":
				{
					cellType = CellBlizzard
					blizzard.direction = DirDown
				}

			}
			areaMap[coord] = cellType
			if cellType == CellBlizzard {
				blizzard.id = blizzardCount
				blizzardMap[blizzardCount] = blizzard
				blizzardCount += 1
			} else if cellType == CellEmpty {
				if y == 0 {
					coordStart = coord
				} else {
					coordEnd = coord
				}
			}

			minX = utils.Min(minX, x)
			maxX = utils.Max(maxX, x)
			minY = utils.Min(minY, y)
			maxY = utils.Max(maxY, y)
		}
	}

	{
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if areaMap[Coord{x, y}] != CellWall {
					minNonWallXForY[y] = x
					break
				}
			}
			for x := maxX; x >= minX; x-- {
				if areaMap[Coord{x, y}] != CellWall {
					maxNonWallXForY[y] = x
					break
				}
			}
		}
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				if areaMap[Coord{x, y}] != CellWall {
					minNonWallYForX[x] = y
					break
				}
			}
			for y := maxY; y >= minY; y-- {
				if areaMap[Coord{x, y}] != CellWall {
					maxNonWallYForX[x] = y
					break
				}
			}
		}
	}

	fmt.Printf("Start: %v\n", coordStart)
	fmt.Printf("End: %v\n", coordEnd)
	fmt.Printf("Blizzard count: %v\n", blizzardCount)

	getStepsMapKey := func(state State) string {
		chunks := []string{}

		chunks = append(chunks, strconv.Itoa(state.coord.x))
		chunks = append(chunks, strconv.Itoa(state.coord.y))
		for _, b := range state.blizzardStates {
			chunks = append(chunks, strconv.Itoa(b.coord.x))
			chunks = append(chunks, strconv.Itoa(b.coord.y))
		}

		return strings.Join(chunks, ";")

	}
	dj := func() {

		pq := make(PriorityQueue, 1)
		initialState := State{
			coord:          coordStart,
			blizzardStates: copyMap(blizzardMap),
		}
		pq[0] = &Item{
			value:    initialState,
			index:    0,
			priority: 0,
		}
		heap.Init(&pq)

		stepsMap := map[string]int{}
		stepsMap[getStepsMapKey((initialState))] = 0

		for pq.Len() > 0 {
			currentItem := heap.Pop(&pq).(*Item)
			currentState := currentItem.value
			// fmt.Printf("%v\n", currentState.coord)

			fmt.Printf("currentState: %v\n", currentState)

			if currentState.coord == coordEnd {
				fmt.Println("Reached end", currentItem.priority)
				return
			}
			currentSteps := stepsMap[getStepsMapKey((currentState))]

			nextBlizzardStates := getNextBlizzardStates(currentState.blizzardStates)
			bCoords := map[Coord]bool{}
			for _, b := range nextBlizzardStates {
				bCoords[b.coord] = true
			}

			for dir, _ := range deltas {
				nextCoord := getNextCoord(currentState.coord, dir)
				if areaMap[nextCoord] == CellWall {
					continue
				}
				if bCoords[nextCoord] {
					continue
				}
				nextState := State{
					coord:          nextCoord,
					blizzardStates: nextBlizzardStates,
				}

				nextItem := &Item{
					value:    nextState,
					priority: currentItem.priority + 1,
				}
				if currentState.coord == coordEnd {
					fmt.Println("Reached end", currentItem.priority)
					return
				}
				nextSteps, exist := stepsMap[getStepsMapKey(nextState)]
				if exist {
					fmt.Println("Exist")
					if currentSteps+1 >= nextSteps {
						break
					}
				}

				stepsMap[getStepsMapKey(nextState)] = currentSteps + 1
				heap.Push(&pq, nextItem)
			}

			// Do nothing
			{
				nextState := State{
					coord:          currentState.coord,
					blizzardStates: nextBlizzardStates,
				}
				nextItem := &Item{
					value:    nextState,
					priority: currentItem.priority + 1,
				}

				nextSteps, exist := stepsMap[getStepsMapKey(nextState)]
				if exist {
					fmt.Println("Exist")
					if currentSteps+1 >= nextSteps {
						break
					}
				}
				stepsMap[getStepsMapKey(nextState)] = currentSteps + 1
				heap.Push(&pq, nextItem)
			}

		}
	}
	dj()
}
