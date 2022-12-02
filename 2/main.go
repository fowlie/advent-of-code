package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	rock     = 1
	paper    = 2
	scissors = 3
	win      = 6
	draw     = 3
	lose     = 0
)

func main() {

	reader, _ := os.Open("input")

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var total1 int
	var total2 int

	for scanner.Scan() {
		line := scanner.Text()
		total1 = total1 + part1GetScore(line)
		total2 = total2 + part2GetScore(line)
	}

	fmt.Println("The total score for part 1 is", total1)
	if total1 != 14297 {
		panic("program corrupt")
	}

	fmt.Println("The total score for part 2 is", total2)
	if total2 != 10498 {
		panic("program corrupt")
	}
}

func part2GetScore(line string) int {
	line = strings.Replace(line, "A", "rock", 1)
	line = strings.Replace(line, "B", "paper", 1)
	line = strings.Replace(line, "C", "scissors", 1)

	line = strings.Replace(line, "X", "lose", 1)
	line = strings.Replace(line, "Y", "draw", 1)
	line = strings.Replace(line, "Z", "win", 1)

	switch line {
	case "rock lose":
		line = strings.Replace(line, "lose", "scissors", 1)
	case "rock draw":
		line = strings.Replace(line, "draw", "rock", 1)
	case "rock win":
		line = strings.Replace(line, "win", "paper", 1)

	case "paper lose":
		line = strings.Replace(line, "lose", "rock", 1)
	case "paper draw":
		line = strings.Replace(line, "draw", "paper", 1)
	case "paper win":
		line = strings.Replace(line, "win", "scissors", 1)

	case "scissors lose":
		line = strings.Replace(line, "lose", "paper", 1)
	case "scissors draw":
		line = strings.Replace(line, "draw", "scissors", 1)
	case "scissors win":
		line = strings.Replace(line, "win", "rock", 1)
	}

	return part1GetScore(line)
}

func part1GetScore(line string) int {
	line = strings.Replace(line, "A", "rock", 1)
	line = strings.Replace(line, "B", "paper", 1)
	line = strings.Replace(line, "C", "scissors", 1)

	line = strings.Replace(line, "X", "rock", 1)
	line = strings.Replace(line, "Y", "paper", 1)
	line = strings.Replace(line, "Z", "scissors", 1)

	switch line {
	case "rock paper":
		return paper + win
	case "rock rock":
		return rock + draw
	case "rock scissors":
		return scissors + lose

	case "paper scissors":
		return scissors + win
	case "paper paper":
		return paper + draw
	case "paper rock":
		return rock + lose

	case "scissors rock":
		return rock + win
	case "scissors scissors":
		return scissors + draw
	case "scissors paper":
		return paper + lose

	default:
		panic("invalid line: " + line)
	}
}
