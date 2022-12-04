package main

import (
	"strconv"
	"strings"
	"utils"
)

type Section struct {
	start int
	end   int
}

func main() {
	var lines = utils.ReadFileToLines("day04.in")
	var completeOverlapCount = 0
	var partialOverlapCount = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		stringSections := strings.Split(line, ",")
		var sections []Section
		for _, sectonString := range stringSections {
			sections = append(sections, parseSectionString(sectonString))
		}
		sectionFirst := sections[0]
		sectionSecond := sections[1]

		// var hasOverlap bool
		if sectionFirst.start <= sectionSecond.start &&
			sectionFirst.end >= sectionSecond.end {
			/*
				1: [            ]
				2:     [   ]
			*/
			completeOverlapCount += 1
			partialOverlapCount += 1
			// hasOverlap = true
		} else if sectionFirst.start >= sectionSecond.start &&
			sectionFirst.end <= sectionSecond.end {
			/*
				1:     [   ]
				2: [            ]
			*/
			completeOverlapCount += 1
			partialOverlapCount += 1
			// hasOverlap = true
		} else if sectionFirst.start <= sectionSecond.start &&
			sectionFirst.end >= sectionSecond.start {
			/*
				1: [            ]
				2:        [         ]
			*/
			partialOverlapCount += 1
			// hasOverlap = true
		} else if sectionFirst.start >= sectionSecond.start &&
			sectionFirst.start <= sectionSecond.end {
			/*
				1:             [            ]
				2:        [         ]
			*/
			partialOverlapCount += 1
			// hasOverlap = true
		}

		// println(line, ":", hasOverlap)

	}
	println("Part 1: ", completeOverlapCount)
	println("Part 2: ", partialOverlapCount)
}

func parseSectionString(sectonString string) Section {
	var sectionIds []int

	for _, sectionIdStr := range strings.Split(sectonString, "-") {
		sectionId, _ := strconv.Atoi(sectionIdStr)
		sectionIds = append(sectionIds, sectionId)
	}

	return Section{start: sectionIds[0], end: sectionIds[1]}
}
