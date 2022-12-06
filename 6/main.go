package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input")
	for i := 0; i < len(data)-4; i++ {
		if uniqueBytes(data[i : i+4]) {
			fmt.Println("Part 1 result is", i+4)
			break
		}
	}


	data, _ = os.ReadFile("input")
	for i := 0; i < len(data)-14; i++ {
		if uniqueBytes(data[i : i+14]) {
			fmt.Println("Part 2 result is", i+14)
			break
		}
	}
}

func uniqueBytes(bytes []byte) bool {
	for i := 0; i < len(bytes); i++ {
    for j := 0; j < len(bytes); j++ {
      if i == j {
        break
      }
      if bytes[i] == bytes[j] {
        return false
      }
    }
	}
	return true
}
