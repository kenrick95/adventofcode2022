package main

import (
	"math"
	"regexp"
	"strconv"
	"utils"
)

type NodeType int
type MathOp int

const (
	Number NodeType = iota + 1
	Math
)
const (
	Add MathOp = iota + 1
	Sub
	Mul
	Div
)

type Node struct {
	name       string
	nodeType   NodeType
	value      int
	op         MathOp
	operandId1 string
	operandId2 string
}

func main() {
	lines := utils.ReadFileToLines("day21.in")
	reMath, _ := regexp.Compile(`([a-z]+): ([a-z]+) (\*|\+|\/|\-) ([a-z]+)`)
	reNumber, _ := regexp.Compile(`([a-z]+): (\d+)`)

	nodeMap := map[string]*Node{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		matchesNumber := reNumber.FindStringSubmatch(line)
		matchesMath := reMath.FindStringSubmatch(line)
		if matchesNumber != nil {
			id := matchesNumber[1]
			num, _ := strconv.Atoi(matchesNumber[2])

			nodeMap[id] = &Node{
				name:       id,
				nodeType:   Number,
				value:      num,
				op:         0,
				operandId1: "",
				operandId2: "",
			}
		} else if matchesMath != nil {
			id := matchesMath[1]
			idOp1 := matchesMath[2]
			rawOp := matchesMath[3]
			idOp2 := matchesMath[4]
			op := Add
			switch rawOp {
			case "+":
				{
					op = Add
				}
			case "-":
				{
					op = Sub
				}
			case "*":
				{
					op = Mul
				}
			case "/":
				{
					op = Div
				}
			}

			nodeMap[id] = &Node{
				name:       id,
				nodeType:   Math,
				value:      0,
				op:         op,
				operandId1: idOp1,
				operandId2: idOp2,
			}
		}
	}

	visMap := map[string]int{}

	var eval func(id string) int
	eval = func(id string) int {
		// println("eval", id)
		node, exist := nodeMap[id]
		if !exist {
			return 0
		}

		nodeValue, nodeVisited := visMap[id]
		if nodeVisited {
			return nodeValue
		}

		if node.nodeType == Number {
			visMap[id] = node.value
			// println("eval", id, "num", node.value)
			return node.value
		}

		op1Value := eval(node.operandId1)
		op2Value := eval(node.operandId2)

		res := op1Value

		switch node.op {
		case Add:
			{
				res += op2Value
			}
		case Sub:
			{
				res -= op2Value
			}
		case Mul:
			{
				res *= op2Value
			}
		case Div:
			{
				res /= op2Value
			}
		}

		visMap[id] = res
		// println("eval", id, "math", node.operandId1, node.operandId2, res)
		return res
	}

	// ansPart1 := eval("root")
	// println("ansPart1", ansPart1)

	{
		// Part 2
		rootNode := nodeMap["root"]
		rootNodeOp1 := rootNode.operandId1
		rootNodeOp2 := rootNode.operandId2
		humnNode := nodeMap["humn"]

		// Observation: rootNodeOp2Val doesn't change when I change humn.value
		// So I can just binary search since rootNodeOp1Val seem to be decreasing with increasing humn.value
		// too low:   3_625_000_000_000
		// too high:  3_750_000_000_000

		rootNodeOp2Val := eval(rootNodeOp2)

		{
			// binary search
			lo := 3_625_000_000_000
			hi := 3_750_000_000_000

			for lo != hi {
				mid := int(math.Ceil(float64(lo+hi) / 2))
				// println("check", "lo", lo, "mid", mid, "hi", hi)

				visMap = map[string]int{}
				humnNode.value = mid

				rootNodeOp1Val := eval(rootNodeOp1)
				diff := rootNodeOp1Val - rootNodeOp2Val
				// println("val", rootNodeOp1Val, diff)

				if diff == 0 {
					println("ansPart2", mid)
					break
				} else if diff < 0 {
					hi = mid - 1
				} else {
					lo = mid
				}
			}
			// println(lo, hi)
		}

	}
}
