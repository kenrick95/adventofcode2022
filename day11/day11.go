package main

import (
	"strconv"
	"strings"
	"utils"
)

type Operation int

const (
	OperationPlus Operation = iota + 1
	OperationMinus
	OperationMultiply
	OperationDivision
)

type OperandType int

const (
	OperandTypeLiteral OperandType = iota + 1
	OperandTypeVariable
)

type Operand struct {
	operandType OperandType
	literal     uint64
}

type Monkey struct {
	items           []*Item
	index           int
	inspectionCount uint64

	operation Operation
	operands  []Operand

	divisibilityTestOperand       uint64
	nextMonkeyIfDivisible         *Monkey
	nextMonkeyIndexIfDivisible    int
	nextMonkeyIfNotDivisible      *Monkey
	nextMonkeyIndexIfNotDivisible int
}

type Item struct {
	worryLevel uint64
	monkey     *Monkey
}

func main() {
	lines := utils.ReadFileToLines("day11.in")
	var PRIME_MODULO uint64 = 1

	// Parse input
	monkeyMap := map[int]*Monkey{}
	itemMap := map[int]*Item{}
	monkeyIndex := 0
	itemIndex := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey") {
			parts := strings.Split(line, " ")
			idx, _ := strconv.Atoi(strings.Split(parts[1], "")[0])
			monkeyIndex = idx

			monkeyMap[monkeyIndex] = &Monkey{
				index:           monkeyIndex,
				items:           []*Item{},
				inspectionCount: 0,
			}
		} else if strings.HasPrefix(line, "Starting items: ") {
			trimmedLine := strings.TrimPrefix(line, "Starting items: ")
			itemStrings := strings.Split(trimmedLine, ", ")
			for _, itemString := range itemStrings {
				worryLevel, _ := strconv.Atoi(itemString)
				itemMap[itemIndex] = &Item{
					worryLevel: uint64(worryLevel),
					monkey:     monkeyMap[monkeyIndex],
				}
				monkeyMap[monkeyIndex].items = append(monkeyMap[monkeyIndex].items, itemMap[itemIndex])
			}
		} else if strings.HasPrefix(line, "Operation: new = ") {
			trimmedLine := strings.TrimPrefix(line, "Operation: new = ")
			parts := strings.Split(trimmedLine, " ")

			operation := OperationPlus
			if parts[1] == "+" {
				operation = OperationPlus
			} else if parts[1] == "*" {
				operation = OperationMultiply
			} else if parts[1] == "-" {
				operation = OperationMinus
			} else if parts[1] == "/" {
				operation = OperationDivision
			}
			var operandOneType OperandType
			var operandOneLiteral int
			if parts[0] == "old" {
				operandOneType = OperandTypeVariable
				operandOneLiteral = 0
			} else {
				operandOneType = OperandTypeLiteral
				operandOneLiteral, _ = strconv.Atoi(parts[0])
			}
			var operandTwoType OperandType
			var operandTwoLiteral int
			if parts[2] == "old" {
				operandTwoType = OperandTypeVariable
				operandTwoLiteral = 0
			} else {
				operandTwoType = OperandTypeLiteral
				operandTwoLiteral, _ = strconv.Atoi(parts[2])
			}
			operandOne := Operand{
				operandType: operandOneType,
				literal:     uint64(operandOneLiteral),
			}
			operandTwo := Operand{
				operandType: operandTwoType,
				literal:     uint64(operandTwoLiteral),
			}

			monkeyMap[monkeyIndex].operands = append(monkeyMap[monkeyIndex].operands, operandOne)
			monkeyMap[monkeyIndex].operands = append(monkeyMap[monkeyIndex].operands, operandTwo)
			monkeyMap[monkeyIndex].operation = operation

		} else if strings.HasPrefix(line, "Test: divisible by ") {
			trimmedLine := strings.TrimPrefix(line, "Test: divisible by ")
			divisibleAmount, _ := strconv.Atoi(trimmedLine)
			monkeyMap[monkeyIndex].divisibilityTestOperand = uint64(divisibleAmount)
			PRIME_MODULO *= uint64(divisibleAmount)

		} else if strings.HasPrefix(line, "If true: throw to monkey ") {
			trimmedLine := strings.TrimPrefix(line, "If true: throw to monkey ")
			nextMonkeyIndex, _ := strconv.Atoi(trimmedLine)
			monkeyMap[monkeyIndex].nextMonkeyIndexIfDivisible = nextMonkeyIndex

		} else if strings.HasPrefix(line, "If false: throw to monkey ") {
			trimmedLine := strings.TrimPrefix(line, "If false: throw to monkey ")
			nextMonkeyIndex, _ := strconv.Atoi(trimmedLine)
			monkeyMap[monkeyIndex].nextMonkeyIndexIfNotDivisible = nextMonkeyIndex

		}
	}
	totalMonkey := len(monkeyMap)

	for i := 0; i < totalMonkey; i++ {
		monkeyMap[i].nextMonkeyIfDivisible = monkeyMap[monkeyMap[i].nextMonkeyIndexIfDivisible]
		monkeyMap[i].nextMonkeyIfNotDivisible = monkeyMap[monkeyMap[i].nextMonkeyIndexIfNotDivisible]
	}

	// Part 1: ROUND_COUNT = 20
	// Part 2: ROUND_COUNT = 10000
	const ROUND_COUNT = 10000
	for round := 1; round <= ROUND_COUNT; round++ {
		// println("Round", round)
		for i := 0; i < totalMonkey; i++ {
			// println("Monkey", i)
			// Inspect items
			items := monkeyMap[i].items
			operands := monkeyMap[i].operands
			operation := monkeyMap[i].operation
			divisibilityTestOperand := monkeyMap[i].divisibilityTestOperand
			for _, item := range items {
				// println("Item: worry level", item.worryLevel)
				newWorryLevel := uint64(0)

				monkeyMap[i].inspectionCount += 1

				// Do operation
				if operation == OperationMultiply {
					if operands[0].operandType == OperandTypeVariable {
						newWorryLevel = item.worryLevel
					} else if operands[0].operandType == OperandTypeLiteral {
						newWorryLevel = operands[0].literal
					}
					if operands[1].operandType == OperandTypeVariable {
						newWorryLevel *= item.worryLevel
					} else if operands[1].operandType == OperandTypeLiteral {
						newWorryLevel *= operands[1].literal
					}
				} else if operation == OperationPlus {
					if operands[0].operandType == OperandTypeVariable {
						newWorryLevel = item.worryLevel
					} else if operands[0].operandType == OperandTypeLiteral {
						newWorryLevel = operands[0].literal
					}
					if operands[1].operandType == OperandTypeVariable {
						newWorryLevel += item.worryLevel
					} else if operands[1].operandType == OperandTypeLiteral {
						newWorryLevel += operands[1].literal
					}
				}
				newWorryLevel = newWorryLevel % PRIME_MODULO

				// Part 1: beforeTestWorryLevel := newWorryLevel / 3
				// Part 2: beforeTestWorryLevel := newWorryLevel
				beforeTestWorryLevel := newWorryLevel / 3
				item.worryLevel = beforeTestWorryLevel
				// println("Item: new worry level", item.worryLevel)

				// Test & throw
				if beforeTestWorryLevel%divisibilityTestOperand == 0 {
					// println("Item thrown to monkey", monkeyMap[i].nextMonkeyIndexIfDivisible)
					monkeyMap[i].nextMonkeyIfDivisible.items = append(monkeyMap[i].nextMonkeyIfDivisible.items, item)
				} else {
					// println("Item thrown to monkey", monkeyMap[i].nextMonkeyIndexIfNotDivisible)
					monkeyMap[i].nextMonkeyIfNotDivisible.items = append(monkeyMap[i].nextMonkeyIfNotDivisible.items, item)
				}
			}
			monkeyMap[i].items = []*Item{}
			// println()
		}
		// println("End round")

		// for i := 0; i < totalMonkey; i++ {
		// 	print("Monkey ", i, ": ")
		// 	for _, item := range monkeyMap[i].items {
		// 		print(item.worryLevel, ", ")
		// 	}
		// 	println("")
		// }

		if round%1000 == 0 || round == 1 || round == 20 {
			println("Round", round)

			for i := 0; i < totalMonkey; i++ {
				println("Monkey ", i, ": ", monkeyMap[i].inspectionCount)
			}
		}

		// println()

		// println()
	}

	var biggest uint64 = 0
	var secondBiggest uint64 = 0
	for i := 0; i < totalMonkey; i++ {
		if monkeyMap[i].inspectionCount > biggest {
			secondBiggest = biggest
			biggest = monkeyMap[i].inspectionCount
		} else if monkeyMap[i].inspectionCount > secondBiggest {
			secondBiggest = monkeyMap[i].inspectionCount
		}
		println("Count", i, monkeyMap[i].inspectionCount)
	}
	println("Monkey business:", biggest*secondBiggest)
}
