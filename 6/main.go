package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input")
  fmt.Println("Part 1 result is", markerPosition(data, 4))
  fmt.Println("Part 2 result is", markerPosition(data, 14))
}

func markerPosition(data []byte, size int) int {
	for i := 0; i < len(data)-size; i++ {
		if uniqueBytes(data[i : i+size]) {
			return i + size
		}
	}
	return -1
}

func uniqueBytes(bytes []byte) bool {
	for i := 0; i < len(bytes); i++ {
		for j := 0; j < len(bytes); j++ {
			if i != j && bytes[i] == bytes[j] {
				return false
			}
		}
	}
	return true
}
