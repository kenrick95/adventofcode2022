package main

import (
	"regexp"
	"sort"
	"strconv"
	"utils"
)

type Coord struct {
	x int
	y int
}

type Sensor struct {
	coord         Coord
	nearestBeacon *Beacon
}
type Beacon struct {
	coord Coord
}

func main() {
	sensorMap := map[Coord]*Sensor{}
	beaconMap := map[Coord]*Beacon{}

	lines := utils.ReadFileToLines("day15.in")

	re, _ := regexp.Compile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		sensorX, _ := strconv.Atoi(matches[1])
		sensorY, _ := strconv.Atoi(matches[2])
		beaconX, _ := strconv.Atoi(matches[3])
		beaconY, _ := strconv.Atoi(matches[4])

		sensorCoord := Coord{
			x: sensorX,
			y: sensorY,
		}
		beaconCoord := Coord{
			x: beaconX,
			y: beaconY,
		}
		beacon := Beacon{
			coord: beaconCoord,
		}
		beaconMap[beaconCoord] = &beacon
		sensor := Sensor{
			coord:         sensorCoord,
			nearestBeacon: &beacon,
		}
		sensorMap[sensorCoord] = &sensor
	}

	beaconYCount := map[int]int{}
	for beaconCoord := range beaconMap {
		curCount, exist := beaconYCount[beaconCoord.y]
		if !exist {
			beaconYCount[beaconCoord.y] = 1
		} else {
			beaconYCount[beaconCoord.y] = curCount + 1
		}
	}

	// sample
	// const MAX_X = 20
	// const MAX_Y = 20
	// real
	const MAX_X = 4000000
	const MAX_Y = 4000000

	getRanges := func(checkY int, isPart2 bool) [][]int {
		// println("checkY", checkY)
		xRangesAtCheckY := [][]int{}
		for sensorCoord, sensor := range sensorMap {
			radius := manhattanDistance(sensorCoord, sensor.nearestBeacon.coord)
			distToCheckY := abs(sensorCoord.y - checkY)
			if distToCheckY <= radius {
				// This sensor influence "checkY" row
				xRadiusAtCheckY := radius - distToCheckY

				xRangeLow := sensorCoord.x - xRadiusAtCheckY
				xRangeHigh := sensorCoord.x + xRadiusAtCheckY

				if isPart2 {
					xRangeLow = max(0, min(MAX_X, xRangeLow))
					xRangeHigh = max(0, min(MAX_X, xRangeHigh))
				}

				// println("Raw range", xRangeLow, xRangeHigh)
				// If beacon exist at these coords, exclude this exact coord from the ranges
				if sensor.nearestBeacon.coord.x == xRangeLow && sensor.nearestBeacon.coord.y == checkY {
					xRangeLow += 1
				}
				if sensor.nearestBeacon.coord.x == xRangeHigh && sensor.nearestBeacon.coord.y == checkY {
					xRangeHigh -= 1
				}
				// println("Raw range (adjusted)", xRangeLow, xRangeHigh)
				if xRangeLow <= xRangeHigh {
					xRangesAtCheckY = append(xRangesAtCheckY, []int{
						xRangeLow,
						xRangeHigh,
					})
				}

			}
		}
		// sort ranges
		sort.Slice(xRangesAtCheckY, func(i int, j int) bool {
			if xRangesAtCheckY[i][0] < xRangesAtCheckY[j][0] {
				return true
			} else if xRangesAtCheckY[i][0] == xRangesAtCheckY[j][0] {
				if xRangesAtCheckY[i][1] < xRangesAtCheckY[j][1] {
					return true
				}
			}
			return false
		})

		return xRangesAtCheckY

	}

	getCoverage := func(checkY int, isPart2 bool) int {
		xRangesAtCheckY := getRanges(checkY, isPart2)

		coverage := 0
		lastX := -1_000_000_000

		for _, ranges := range xRangesAtCheckY {
			// println("Range: ", ranges[0], "-", ranges[1], ", lastX: ", lastX)

			if lastX < ranges[0] {
				//  *
				//      [    ]
				coverage += ranges[1] - ranges[0] + 1
				// println("Case 1:", ranges[1]-ranges[0]+1)
				lastX = ranges[1]
			} else if ranges[0] <= lastX && lastX < ranges[1] {
				//             *
				//      [       ]
				coverage += ranges[1] - lastX
				// println("Case 2:", ranges[1]-lastX)
				lastX = ranges[1]
			}
		}
		return coverage

	}

	// sample
	// println("Ans part 1: ", getCoverage(10, false))
	// real
	println("Ans part 1: ", getCoverage(2000000, false))

	getTuningFrequency := func() int {
		susX := -1
		susY := -1

		for y := 0; y < MAX_Y; y++ {
			coverage := getCoverage(y, true)
			if coverage < MAX_Y+1 {
				if beaconYCount[y] == 0 {
					// println("sus y", y, coverage)
					susY = y
					// Find x!
					xRanges := getRanges(y, true)
					lastX := -1_000_000_000
					for _, ranges := range xRanges {
						if lastX < ranges[0] {
							//  *
							//      [    ]
							if lastX != -1_000_000_000 {
								// println("sus x", lastX, ranges[0])
								susX = lastX + 1
								break
							}
							lastX = ranges[1]
						} else if ranges[0] <= lastX && lastX < ranges[1] {
							//             *
							//      [       ]
							lastX = ranges[1]
						}
					}
					break
				}
			}
		}
		println("sus coord", susX, susY)
		return susX*4000000 + susY
	}

	println("Ans part 2: ", getTuningFrequency())

}

func manhattanDistance(one Coord, two Coord) int {
	return abs(one.x-two.x) + abs(one.y-two.y)
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func min(one int, two int) int {
	if one < two {
		return one
	}
	return two
}
func max(one int, two int) int {
	if one > two {
		return one
	}
	return two
}
