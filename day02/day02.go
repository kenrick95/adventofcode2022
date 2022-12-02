package main

import (
	"log"
	"os"
	"strings"
)

// https://gosamples.dev/enum/
type Outcome int

const (
	Draw Outcome = iota + 1
	Win
	Lose
)

func main() {
	content, err := os.ReadFile("day02.in")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\n")
	var totalScorePart1 int = 0
	var totalScorePart2 int = 0
	for _, line := range lines {
		var cleanedLine string = strings.TrimSpace(line)
		if cleanedLine == "" {
			continue
		}
		var chars []string = strings.Split(cleanedLine, " ")
		totalScorePart1 += roundScorePart1(chars[0], chars[1])
		totalScorePart2 += roundScorePart2(chars[0], chars[1])
	}
	println("totalScorePart1", totalScorePart1)
	println("totalScorePart2", totalScorePart2)

}

func roundScorePart2(opponentShape string, myOutcomeString string) int {
	var score = 0
	var myOutcome Outcome
	var myShape string

	switch myOutcomeString {
	case "X": // I need to lose
		{
			myOutcome = Lose
		}
	case "Y": // I need to draw
		{
			myOutcome = Draw
		}
	case "Z": // I need to win
		{
			myOutcome = Win
		}
	}

	switch myOutcome {
	case Win:
		{
			score += 6
			switch opponentShape {
			case "A": // Rock
				{
					myShape = "Y" // Paper
				}
			case "B": // Paper
				{
					myShape = "Z" // Scissors
				}
			case "C": // Scissors
				{
					myShape = "X" // Rock
				}
			}
		}
	case Draw:
		{
			score += 3
			switch opponentShape {
			case "A": // Rock
				{
					myShape = "X" // Rock
				}
			case "B": // Paper
				{
					myShape = "Y" // Paper
				}
			case "C": // Scissors
				{
					myShape = "Z" // Scissors
				}
			}
		}
	case Lose:
		{
			score += 0
			switch opponentShape {
			case "A": // Rock
				{
					myShape = "Z" // Scissors
				}
			case "B": // Paper
				{
					myShape = "X" // Rock
				}
			case "C": // Scissors
				{
					myShape = "Y" // Paper
				}
			}
		}
	}

	switch myShape {
	case "X": // Rock
		{
			score += 1
		}
	case "Y": // Paper
		{
			score += 2
		}
	case "Z": // Scissors
		{
			score += 3
		}
	}

	return score
}

func roundScorePart1(opponentShape string, myShape string) int {
	var score = 0

	switch myShape {
	case "X": // Rock
		{
			score += 1
		}
	case "Y": // Paper
		{
			score += 2
		}
	case "Z": // Scissors
		{
			score += 3
		}
	}

	var myOutcome Outcome = Draw

	switch opponentShape {
	case "A": // Rock
		{
			switch myShape {
			case "X": // Rock
				{
					myOutcome = Draw
				}
			case "Y": // Paper
				{
					myOutcome = Win
				}
			case "Z": // Scissors
				{
					myOutcome = Lose
				}
			}
		}
	case "B": // Paper
		{
			switch myShape {
			case "X": // Rock
				{
					myOutcome = Lose
				}
			case "Y": // Paper
				{
					myOutcome = Draw
				}
			case "Z": // Scissors
				{
					myOutcome = Win
				}
			}
		}
	case "C": // Scissors
		{
			switch myShape {
			case "X": // Rock
				{
					myOutcome = Win
				}
			case "Y": // Paper
				{
					myOutcome = Lose
				}
			case "Z": // Scissors
				{
					myOutcome = Draw
				}
			}
		}
	}

	switch myOutcome {
	case Win:
		{
			score += 6
		}
	case Draw:
		{

			score += 3
		}
	case Lose:
		{

			score += 0
		}
	}

	return score
}
