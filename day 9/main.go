package main

func main() {
	input := readInput("./in.txt")
	intInput := stringSliceToIntSlice(input)
	invalidNum := part1(intInput, 25)
	println("Part 1: ", invalidNum)
	println("Part 2: ", part2(intInput, invalidNum))
}

func part1(data []int, preambleLen int) int {
	preamble := []int{}
	preambleMap := map[int]bool{}

	for _, val := range data {
		if len(preamble) > preambleLen {
			dropped := preamble[0]
			preamble = preamble[1:]
			delete(preambleMap, dropped)
			if !findSum(val, preamble, preambleMap) {
				return val
			}
		}
		preamble = append(preamble, val)
		preambleMap[val] = true
	}
	panic("Should have found one number that is not a sum")
}

func findSum(num int, preamble []int, preambleMap map[int]bool) bool {
	// num = candidate + diff
	for _, candidate := range preamble {
		diff := num - candidate
		if _, diffPresent := preambleMap[diff]; diffPresent && diff != num {
			return true
		}
	}
	return false
}

func part2(data []int, target int) int {
	for i := 0; i < len(data); i++ {
		sum := data[i]
		min := data[i]
		max := data[i]
		for j := 1; i+j < len(data); j++ {
			nextVal := data[i+j]
			if nextVal < min {
				min = nextVal
			}
			if nextVal > max {
				max = nextVal
			}
			sum += nextVal
			if sum == target {
				return min + max
			}
		}
	}
	panic("Should have already returned answer")
}
