package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader, _ := os.Open("example")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
    first := line[0:len(line)/2]
    second := line[len(line)/2:]
    for _, c := range first {
      if strings.Contains(second, string(c)) {
        fmt.Println("found match", string(c))
      }
    }
  }
}
