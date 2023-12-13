package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Part struct {
	number   int
	position int
}

type ParsedLine struct {
	parts   []Part
	symbols []int
}

type Gear struct {
	part1, part2 Part
}

func FindAdjacentParts(position int, parts []Part) []Part {
	adjacents := make([]Part, 0)
	for _, p := range parts {
		digits := strconv.Itoa(p.number)

		for i := range digits {
			if p.position+i == position || p.position+i == position-1 || p.position+i == position+1 {
				adjacents = append(adjacents, p)
				break
			}
		}
	}

	return adjacents
}

func FindGears(input []ParsedLine) []Gear {
	gears := make([]Gear, 0)
	for i := range input {
		for _, symbol := range input[i].symbols {

			adjacents := FindAdjacentParts(symbol, input[i].parts)
			if i > 0 {
				adjacents = append(adjacents, FindAdjacentParts(symbol, input[i-1].parts)...)
			}
			if i < len(input)-1 {
				adjacents = append(adjacents, FindAdjacentParts(symbol, input[i+1].parts)...)
			}

			if len(adjacents) == 2 {
				gears = append(gears, Gear{adjacents[0], adjacents[1]})
			}
		}
	}
	return gears
}

func ReadFile(path string) []string {
	lines := make([]string, 0)

	readFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func ParseInput(lines []string) []ParsedLine {
	data := make([]ParsedLine, 0)
	for _, line := range lines {
		pl := ParsedLine{
			parts:   FindPartNumbers(line),
			symbols: FindSymbols(line),
		}
		data = append(data, pl)
	}
	return data
}

func main() {
	data := ParseInput(ReadFile("input"))
	gears := FindGears(data)
	answer := 0
	for _, gear := range gears {
		answer += gear.part1.number * gear.part2.number
	}
	// Correct answer is 80703636
	fmt.Printf("The answer is %d\n", answer)
}

func FindSymbols(s string) []int {
	result := make([]int, 0)
	for i, r := range s {
		if r == '*' {
			result = append(result, i)
		}
	}
	return result
}

func FindPartNumbers(s string) []Part {
	matches := []Part{}

	parsed_number := ""
	number := 0

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit([]rune(s)[i]) {
			parsed_number = fmt.Sprintf("%s%s", parsed_number, string([]rune(s)[i]))
			continue
		}
		if len(parsed_number) > 0 {
			var err error
			number, err = strconv.Atoi(parsed_number)
			if err != nil {
				panic(err)
			}
			matches = append(matches, Part{
				position: i - len(parsed_number),
				number:   number,
			})
			parsed_number = ""
		}
	}

	if len(parsed_number) > 0 {
		var err error
		number, err = strconv.Atoi(parsed_number)
		if err != nil {
			panic(err)
		}
		matches = append(matches, Part{
			position: len(s) - len(parsed_number),
			number:   number,
		})
		parsed_number = ""

	}

	return matches
}

func CalculateRatios(gears []Gear) []int {
	ratios := make([]int, len(gears))
	for _, gear := range gears {
		ratios = append(ratios, gear.part1.number*gear.part2.number)
	}
	return ratios
}
