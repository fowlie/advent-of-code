package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	AIR  int = 0
	ROCK int = 1
	SAND int = 2
)

var (
	debug       bool = false
	leftOffset int  = 0
	rightOffset int  = 600
	level       [][]int
	unitsOfSand int = 0
)

func enableDebug() {
	debug = true
}

func createLevel(width, height int) {
	leftOffset = 0
	unitsOfSand = 0
	level = make([][]int, height)
	for i := range level {
		level[i] = make([]int, width)
	}
}

func addSand(x int) {
	for y := 0; y < len(level)-1; y++ {
		// check below for air
		if level[y+1][x] == AIR {
			continue
		}

		// check below on the sides
		if level[y+1][x-1] == AIR { // check for air below to the left
			x = x - 1
		} else if level[y+1][x+1] == AIR {
			x = x + 1
		}

		// check below for rock
		if level[y+1][x] == ROCK {
			level[y][x] = SAND
			unitsOfSand = unitsOfSand + 1
			break
		}

		// check below for sand
		if level[y+1][x] == SAND {
			level[y][x] = SAND
		}
	}
}

func printIt() {
	if debug {
		fmt.Println("Printing the map:")
		for _, row := range level {
			if len(row[leftOffset:rightOffset]) > 0 {
				for _, item := range row[leftOffset:rightOffset] {
					switch item {
					case AIR:
						fmt.Print(".")
					case ROCK:
						fmt.Print("#")
					case SAND:
						fmt.Print("o")
					}
				}
				fmt.Print("\n")
			}
		}
	}
}

func addLine(x1, y1, x2, y2 int) {
	if y1 == y2 {
		for i := min(x1, x2); i <= max(x1, x2); i++ {
			level[y1][i] = ROCK
		}
	} else {
		for i := min(y1, y2); i <= max(y1, y2); i++ {
			level[i][x1] = ROCK
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func countSandUnits() int {
	return unitsOfSand
}

func main() {
	reader, _ := os.Open("input")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	createLevel(800, 200)

  maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " -> ")
		for i := 0; i < len(elements)-1; i++ {
			x1, _ := strconv.Atoi(strings.Split(elements[i], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(elements[i], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(elements[i+1], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(elements[i+1], ",")[1])
			addLine(x1, y1, x2, y2)
      maxY = max(maxY, y1)
      maxY = max(maxY, y2)
		}
	}

  // add 1.000 units of sand
	for range [1000]int{} {
		addSand(500)
  }

	// enableDebug()
	leftOffset = 400
  rightOffset = 600
	printIt()

	fmt.Println("Number of sand units are", countSandUnits())

  // part 2 - add 50.000 units of sand
  addLine(0, maxY + 2, 799, maxY + 2)
	for {
		addSand(500)
    if level[0][500] == SAND {
      break
    }
	}
	printIt()

	fmt.Println("Number of sand units are", countSandUnits())
}
