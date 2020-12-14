package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./in.txt")
	part1(input)
	part2(input)
}

func parseMask(mask string) map[int]int32 {
	parsed := map[int]int32{}
	if len(mask) != 36 {
		panic("Mask of invalid length")
	}
	for idx, override := range mask {
		if override != 'X' {
			parsed[idx] = override
		}
	}
	return parsed
}

func parseAssignment(setPos, setVal string) (uint64, uint64) {
	position, err := strconv.ParseUint(setPos[4:len(setPos)-1], 10, 36)
	if err != nil {
		panic(err)
	}

	value, err := strconv.ParseUint(setVal, 10, 36)
	if err != nil {
		panic(err)
	}

	return position, value
}

func applyMask(mask map[int]int32, val uint64) uint64 {
	convertedVal := ""
	binRepr := fmt.Sprintf("%036s", strconv.FormatUint(val, 2))
	for idx, val := range binRepr {
		if newVal, ok := mask[idx]; ok {
			convertedVal += string(newVal)
		} else {
			convertedVal += string(val)
		}
	}
	finalVal, err := strconv.ParseUint(convertedVal, 2, 36)
	if err != nil {
		panic(err)
	}
	return finalVal
}

func part1(data []string) {
	memory := map[uint64]uint64{}
	// which bits to override: position from beginning starting from 0 -> new val
	var mask map[int]int32
	for _, line := range data {
		lineSplit := strings.Split(line, " = ")
		assignee := lineSplit[0]
		assignedData := lineSplit[1]
		if assignee == "mask" {
			mask = parseMask(assignedData)
		} else {
			memoryPos, memoryVal := parseAssignment(assignee, assignedData)
			memory[memoryPos] = applyMask(mask, memoryVal)
			//fmt.Printf("Set mem[%d] to %d instead of %d\n", memoryPos, memory[memoryPos], memoryVal)
		}
	}

	sum := uint64(0)
	for _, val := range memory {
		sum += val
	}
	println("Part 1:", sum)
}

func generateMasks(mask string, idx int, ) []string {
	if idx < len(mask) {
		if mask[idx] == 'X' {
			newMask := mask[0:idx] + string('2') + mask[idx+1:]
			arr1 := generateMasks(newMask, idx+1)
			newMask = mask[0:idx] + string('1') + mask[idx+1:]
			arr2 := generateMasks(newMask, idx+1)
			return append(arr1, arr2...)
		} else {
			return generateMasks(mask, idx+1)
		}
	} else {
		return []string{mask}
	}
}

func applyMask2(mask map[int]int32, val uint64) uint64 {
	convertedVal := ""
	binRepr := fmt.Sprintf("%036s", strconv.FormatUint(val, 2))
	for idx, val := range binRepr {
		if newVal, ok := mask[idx]; ok && newVal == '1' {
			convertedVal += string(newVal)
		} else if newVal == '2' {
			convertedVal += string('0')
		} else {
			convertedVal += string(val)
		}
	}
	finalVal, err := strconv.ParseUint(convertedVal, 2, 36)
	if err != nil {
		panic(err)
	}
	return finalVal
}

func part2(data []string) {
	memory := map[uint64]uint64{}
	// which bits to override: position from beginning starting from 0 -> new val
	var mask string
	for _, line := range data {
		lineSplit := strings.Split(line, " = ")
		assignee := lineSplit[0]
		assignedData := lineSplit[1]
		if assignee == "mask" {
			mask = assignedData
		} else {
			memoryPos, memoryVal := parseAssignment(assignee, assignedData)
			for _, curMask := range generateMasks(mask, 0) {
				newPos := applyMask2(parseMask(curMask), memoryPos)
				memory[newPos] = memoryVal
				//fmt.Printf("Set mem[%d] instead of instead of %d\n", newPos, memoryPos)
			}
		}
	}
	sum := uint64(0)
	for _, val := range memory {
		sum += val
	}
	println("Part 2:", sum)

}
