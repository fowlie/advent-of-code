package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader, _ := os.Open("input")

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var elves []int

	var calories int
	var number int
	for scanner.Scan() {
		line := scanner.Text()
		number, _ = strconv.Atoi(line)

		if len(line) > 0 {
			calories = calories + number
		} else {
			elves = append(elves, calories)
			calories = 0
			number = 0
		}
	}

	// asume no empty line at end of file
	elves = append(elves, calories+number)

	sort.Ints(elves)
  fmt.Println("The elf with the most calories has", elves[len(elves)-1], "calories.")

}
