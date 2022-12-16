package main

import (
	"regexp"
	"strconv"
	"strings"
	"utils"
)

type Node struct {
	name         string
	flowRate     int
	destinations []*Node
}

type Dp struct {
	nodeName    string
	timeLeft    int
	openedNodes int
}

func main() {
	lines := utils.ReadFileToLines("day16.in")
	re, _ := regexp.Compile(`(?i)Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z ,]+)`)

	nodeMap := map[string]*Node{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		// println("line", line)
		matches := re.FindStringSubmatch(line)
		nodeName := matches[1]
		nodeFlowRate, _ := strconv.Atoi(matches[2])
		nodeDestinations := strings.Split(matches[3], ", ")

		node, nodeExist := nodeMap[nodeName]
		if nodeExist {
			node.flowRate = nodeFlowRate
		} else {
			nodeMap[nodeName] = &Node{
				name:     nodeName,
				flowRate: nodeFlowRate,
			}
		}

		for _, nodeDestinationName := range nodeDestinations {
			_, nodeDestExist := nodeMap[nodeDestinationName]
			if !nodeDestExist {
				nodeMap[nodeDestinationName] = &Node{
					name: nodeDestinationName,
				}
			}
			nodeMap[nodeName].destinations = append(nodeMap[nodeName].destinations, nodeMap[nodeDestinationName])
		}

	}

	nodeIndexMapToName := map[int]string{}
	nodeNameMapToIndex := map[string]int{}
	i := 0
	for nodeName := range nodeMap {
		nodeIndexMapToName[i] = nodeName
		nodeNameMapToIndex[nodeName] = i
		i++
	}

	dpPart1 := map[Dp]int{}
	visitedMap := map[int]map[string]bool{}

	/*
		openedNodes = bitmask of opened node index
		Return total eventual pressure
	*/
	var traverse func(nodeName string, openedNodes int, timeLeft int) int
	traverse = func(nodeName string, openedNodes int, timeLeft int) int {

		// println("t", nodeName, openedNodes, timeLeft)
		if timeLeft == 0 {
			return 0
		}
		if timeLeft < 0 {
			return -1_000_000_000
		}

		dpValue, dpExist := dpPart1[Dp{nodeName, timeLeft, openedNodes}]
		if dpExist {
			return dpValue
		}

		maxReturn := 0
		node := nodeMap[nodeName]
		for _, nodeDest := range node.destinations {
			// Go and do nothing
			if timeLeft >= 1 {
				maxReturn = max(maxReturn, traverse(nodeDest.name, openedNodes, timeLeft-1))
			}

			// Go and open tap
			// Makes sense only when it has positive flow rate
			if timeLeft >= 2 && nodeDest.flowRate > 0 {
				mask := 1 << nodeNameMapToIndex[nodeDest.name]
				newOpenedNodes := openedNodes | mask

				_, visExist := visitedMap[newOpenedNodes]
				if !visExist {
					visitedMap[newOpenedNodes] = map[string]bool{}
				}
				if !visitedMap[newOpenedNodes][nodeDest.name] {
					visitedMap[newOpenedNodes][nodeDest.name] = true
					maxReturn = max(maxReturn, traverse(nodeDest.name, openedNodes|mask, timeLeft-2)+(timeLeft-2)*nodeDest.flowRate)
					visitedMap[newOpenedNodes][nodeDest.name] = false
				}
			}
		}
		dpPart1[Dp{nodeName, openedNodes, timeLeft}] = maxReturn
		return maxReturn
	}
	visitedMap[0] = map[string]bool{}
	visitedMap[0]["AA"] = true
	ansPart1 := traverse("AA", 0, 30)
	println("Ans part 1", ansPart1)

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
