package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	path     string
	dir      bool
	size     int
	parent   *Node
	children []*Node
}

func (n *Node) addDir(name string) {
	n.children = append(n.children, &Node{name, true, 0, n, nil})
}

func (n *Node) addFile(name string, size int) {
	n.children = append(n.children, &Node{name, false, size, n, nil})
}

func (n *Node) contains(path string) bool {
	for _, child := range n.children {
		if child.path == path {
			return true
		}
	}
	return false
}

func cd(path string, head *Node) *Node {
	if path == ".." && head.parent != nil {
		head = head.parent
	} else if head.contains(path) {
		for _, child := range head.children {
			if child.path == path {
				head = child
			}
		}
	}
	return head
}

func calcSize(n *Node) {
	for _, child := range n.children {
		calcSize(child)
	}
	if n.parent != nil {
		n.parent.size = n.parent.size + n.size
	}
}

func getDirSizes(n *Node) {
	if n.dir {
		dirSizes = append(dirSizes, n.size)
	}
	for _, child := range n.children {
		getDirSizes(child)
	}
}

var (
	root     = Node{"/", true, 0, nil, nil}
	dirSizes = make([]int, 0)
)

func main() {
	reader, _ := os.Open("input")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	head := &root

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "$ cd") {
			head = cd(line[5:], head)
		} else if !strings.Contains(line, "$ ls") {
			// parse ls command
			name := strings.Split(line, " ")[1]
			meta := strings.Split(line, " ")[0]
			if meta == "dir" {
				head.addDir(name)
			} else {
				size, _ := strconv.Atoi(meta)
				head.addFile(name, size)
			}
		}
	}

	calcSize(&root)
	//print(root, "")
	getDirSizes(&root)
	sort.Ints(dirSizes)

	total := 0
	for _, size := range dirSizes {
		if size <= 100000 {
			total = total + size
		}
	}

	fmt.Println("Sum of dir size of at most 100000 is", total)

	// part 2
	needed := 30000000 - (70000000 - root.size)
	for i := 1; i < len(dirSizes); i++ {
		if dirSizes[i] > needed {
			fmt.Println("Delete the folder of size", dirSizes[i])
			break
		}
	}
}

func print(n Node, indent string) {
	fmt.Print(indent + "- " + n.path)
	if n.dir {
		fmt.Printf(" (dir, size=%d)\n", n.size)
	} else {
		fmt.Printf(" (file, size=%d)\n", n.size)
	}
	for _, child := range n.children {
		print(*child, indent+"  ")
	}
}
