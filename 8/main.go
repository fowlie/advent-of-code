package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	data   [][]int
	width  int
	height int
)

func main() {
	bytes, _ := os.ReadFile("input")
	lines := strings.Split(string(bytes), "\n")

	width = strings.Index(string(bytes), "\n")
	height = len(lines) - 1

	// parse input into nested array
	data = make([][]int, width)
	for _, line := range lines {
		for j, c := range line {
			val, _ := strconv.Atoi(string(c))
			data[j] = append(data[j], val)
		}
	}

	score := 0
	scenic := 0
	maxScenic := 0
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if visibleFromBottom(x, y) || visibleFromTop(x, y) || visibleFromLeft(x, y) || visibleFromRight(x, y) {
				score = score + 1
			}
			scenic = scenicScore(x, y)
			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
	}
	edges := width*4 - 4
	fmt.Printf("With %d trees visible on the edge and %d visible in the interior, a total of %d trees are visible.\n", edges, score, edges+score)

	fmt.Println("Max scenic score is", maxScenic)
}

func visibleFromRight(x int, y int) bool {
	for i := x + 1; i < width; i++ {
		if data[i][y] >= data[x][y] {
			return false
		}
	}
	return true
}

func visibleFromBottom(x int, y int) bool {
	for i := y + 1; i < height; i++ {
		if data[x][i] >= data[x][y] {
			return false
		}
	}
	return true
}

func visibleFromLeft(x int, y int) bool {
	for i := x - 1; i >= 0; i-- {
		if data[i][y] >= data[x][y] {
			return false
		}
	}
	return true
}

func visibleFromTop(x int, y int) bool {
	for i := y - 1; i >= 0; i-- {
		if data[x][i] >= data[x][y] {
			return false
		}
	}
	return true
}

func scenicScore(x int, y int) int {
	up := 0
	for i := y - 1; i >= 0; i-- {
		up = up + 1
		if data[x][i] >= data[x][y] {
			break
		}
	}

	down := 0
	for i := y + 1; i < height; i++ {
		down = down + 1
		if data[x][i] >= data[x][y] {
			break
		}
	}

	right := 0
	for i := x + 1; i < height; i++ {
		right = right + 1
		if data[i][y] >= data[x][y] {
			break
		}
	}

	left := 0
	for i := x - 1; i >= 0; i-- {
		left = left + 1
		if data[i][y] >= data[x][y] {
			break
		}
	}

	return up * left * right * down
}
