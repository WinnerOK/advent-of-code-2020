package main

import "fmt"

const (
	emptySeat    = 1
	floor        = 0
	occupiedSeat = -1
)

const (
	emptySeatChar    = 'L'
	floorChar        = '.'
	occupiedSeatChar = '#'
)

func main() {
	input := readInput("./in.txt")
	part1(input)
	part2(input)
}

func generateSeatMap(data []string) [][]int {
	seatMap := make([][]int, len(data))
	for i := range seatMap {
		seatMap[i] = make([]int, len(data[0]))
	}

	for rowIdx, row := range data {
		for colIdx, el := range row {
			seatMap[rowIdx][colIdx] = encodeEl(el)
		}
	}
	return seatMap
}

func countOccupied(seatMap [][]int) int {
	occupied := 0
	for _, row := range seatMap {
		for _, el := range row {
			if el == occupiedSeat {
				occupied += 1
			}
		}
	}
	return occupied
}

func encodeEl(el int32) int {
	switch el {
	case emptySeatChar:
		return emptySeat
	case floorChar:
		return floor
	case occupiedSeatChar:
		return occupiedSeat
	default:
		panic("Unknown element")
	}
}

func countAdjacentOccupied(seatMap [][]int, row, col int) int {
	adjacentOccupied := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if !(dx == 0 && dy == 0) {
				rowIdx := row + dy
				colIdx := col + dx
				if 0 <= rowIdx && rowIdx < len(seatMap) &&
					0 <= colIdx && colIdx < len(seatMap[0]) {
					if seatMap[rowIdx][colIdx] == occupiedSeat {
						adjacentOccupied += 1
					}
				}
			}
		}
	}
	return adjacentOccupied
}

func shouldChange(cur int, adjOccupied int, freeIfSeeOccupied int) bool {
	if cur == emptySeat && adjOccupied == 0 {
		return true
	} else if cur == occupiedSeat && adjOccupied >= freeIfSeeOccupied {
		return true
	} else {
		return false
	}
}

func printSeatMap(seatMap [][]int) {
	for _, row := range seatMap {
		for _, el := range row {
			var toPrint int32
			switch el {
			case occupiedSeat:
				toPrint = occupiedSeatChar
			case floor:
				toPrint = floorChar
			case emptySeat:
				toPrint = emptySeatChar
			}
			fmt.Printf("%c", toPrint)
		}
		println()
	}
	println()
}

func simulate(seatMap [][]int, freeSeatIfSeeOccupied int, seatCountFunc func([][]int, int, int) int) [][]int {
	changed := true

	for changed {
		changed = false
		seatMapCopy := make([][]int, len(seatMap))
		for i := range seatMap {
			seatMapCopy[i] = make([]int, len(seatMap[i]))
			copy(seatMapCopy[i], seatMap[i])
		}
		for rowIdx := 0; rowIdx < len(seatMap); rowIdx++ {
			for colIdx := 0; colIdx < len(seatMap[0]); colIdx++ {
				if curPlace := seatMap[rowIdx][colIdx]; curPlace != floor {
					adj := seatCountFunc(seatMap, rowIdx, colIdx)
					if shouldChange(curPlace, adj, freeSeatIfSeeOccupied) {
						changed = true
						seatMapCopy[rowIdx][colIdx] *= -1
					}
				}
			}
		}
		seatMap = seatMapCopy
		//printSeatMap(seatMap)
	}
	return seatMap
}

func part1(data []string) {
	seatMap := generateSeatMap(data)
	seatMap = simulate(seatMap, 4, countAdjacentOccupied)
	println("Part 1:", countOccupied(seatMap))
}

func countDirectionalOccupied(seatMap [][]int, row, col int) int {
	directionalOccupied := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if !(dx == 0 && dy == 0) {
				rowIdx := row + dy
				colIdx := col + dx
				for 0 <= rowIdx && rowIdx < len(seatMap) &&
					0 <= colIdx && colIdx < len(seatMap[0]) {
					if seatMap[rowIdx][colIdx] != floor {
						if seatMap[rowIdx][colIdx] == occupiedSeat {
							directionalOccupied += 1
						}
						break
					}
					rowIdx += dy
					colIdx += dx
				}

			}
		}
	}
	return directionalOccupied
}

func part2(data []string) {
	seatMap := generateSeatMap(data)
	seatMap = simulate(seatMap, 5, countDirectionalOccupied)
	println("Part 2:", countOccupied(seatMap))
}
