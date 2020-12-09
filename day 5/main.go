package main

import (
	"fmt"
)

const (
	lowerHalf       = 'F'
	upperHalf       = 'B'
	upperHalfColumn = 'R'
	lowerHalfColumn = 'L'
)

const (
	rowLow  = 0
	rowHigh = 127
	rowLen  = rowHigh - rowLow + 1

	columnLow  = 0
	columnHigh = 7
	columnLen  = columnHigh - columnLow + 1
)

func getByGroup(splitDescription string, low, high int, lowSymbol, highSymbol int32) int {
	for _, group := range splitDescription {
		if group == highSymbol {
			low = (low + high + 1) / 2
		} else if group == lowSymbol {
			high = (low+high+1)/2 - 1
		}
	}
	return low
}

func getSeatID(seatDescription string) (int, int, int) {
	if len(seatDescription) != 10 {
		panic(fmt.Sprintf("Invalid seat description length: %d", len(seatDescription)))
	}
	row := getByGroup(seatDescription[:7], rowLow, rowHigh, lowerHalf, upperHalf)
	column := getByGroup(seatDescription[7:], columnLow, columnHigh, lowerHalfColumn, upperHalfColumn)

	//println(row, column)

	return row, column, row*8 + column
}

func part1(data []string) {
	maxSeatID := 0
	for _, seatDescription := range data {
		if _, _, seatID := getSeatID(seatDescription); seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Printf("Part 1: %d\n", maxSeatID)
}

func part2(data []string) {
	seatmap := [rowLen * columnLen]int{}
	for _, seatDescription := range data {
		row, column, seatID := getSeatID(seatDescription)
		seatmap[row * columnLen + column] = seatID
	}

	lastEmpty := -1
	for idx, id := range seatmap {
		if id == 0 {
			if idx != lastEmpty+1 {
				println()
			}
			print(idx, " ")
			lastEmpty = idx
		}
	}

	fmt.Printf("Answer to part 2 is middle number")
}

func main() {
	input := readInput("./in.txt")
	part1(input)
	println("---")
	part2(input)
}
