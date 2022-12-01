package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// content, err := os.ReadFile("day01-sample.in")
	content, err := os.ReadFile("day01.in")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\n")
	var elfCalories []int
	var caloriesForCurrentElf int = 0
	for _, line := range lines {
		var cleanedLine string = strings.TrimSpace(line)
		if cleanedLine == "" {
			// break for next elf, so we compare
			elfCalories = append(elfCalories, caloriesForCurrentElf)
			caloriesForCurrentElf = 0
			continue
		}
		var number, _ = strconv.Atoi(cleanedLine)
		caloriesForCurrentElf += number
	}
	// check for final elf
	if caloriesForCurrentElf > 0 {
		elfCalories = append(elfCalories, caloriesForCurrentElf)
	}

	sort.Ints(elfCalories)

	var answerPart1 int = elfCalories[len(elfCalories)-1]
	fmt.Println(answerPart1)

	var answerPart2 int = elfCalories[len(elfCalories)-1] + elfCalories[len(elfCalories)-2] + elfCalories[len(elfCalories)-3]
	fmt.Println(answerPart2)
}
