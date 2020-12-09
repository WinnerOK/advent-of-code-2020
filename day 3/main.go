package main

import "fmt"

const (
	tree = '#'
)

func part1(data []string, dx, dy int) int {
	xPos := 0
	yPos := 0
	xMax := len(data[0])
	treeCnt := 0
	for yPos < len(data) {
		if data[yPos][xPos] == tree {
			treeCnt += 1
		}
		xPos = (xPos + dx) % xMax
		yPos = yPos + dy
	}
	println(treeCnt)
	return treeCnt
}

type slope struct {
	dx, dy int
}

func part2(data []string) {
	var slopes = []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	treesFound := []int{}
	for _, s := range slopes {
		treesFound = append(treesFound, part1(data, s.dx, s.dy))
	}
	answer := 1
	for _, f := range treesFound {
		answer *= f
	}
	fmt.Printf("Multiplied: %d\n", answer)

}

func main() {
	input := readInput("./in.txt")
	part1(input, 3, 1)
	println("------------------")
	part2(input)
}
