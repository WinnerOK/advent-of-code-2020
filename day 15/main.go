package main

import (
	"errors"
	"strings"
)

func main() {
	input := readInput("./in.txt")
	data := stringSliceToIntSlice(strings.Split(input[0], ","))
	println("Part 1:", solve(data, 2020))
	println("Part 2:", solve(data, 30000000))
}

type Occurrence struct {
	last, prelast, sawTimes int
}

func (o *Occurrence) add(position int) error {
	if position < o.last {
		return errors.New("new position shouldn't be greater that last")
	}

	o.prelast = o.last
	o.last = position
	o.sawTimes += 1
	return nil
}

func (o Occurrence) diff() int {
	return o.last - o.prelast
}

func (o Occurrence) didSeeTwice() bool {
	return o.sawTimes >= 2
}

func addMapEntry(m map[int]*Occurrence, entry, number int) map[int]*Occurrence {
	if val, ok := m[entry]; ok {
		err := val.add(number)
		if err != nil {
			panic(err)
		}
	} else {
		newOccurrence := Occurrence{last: number, sawTimes: 1}
		m[entry] = &newOccurrence
	}
	return m
}

func solve(data []int, terminationTurn int) int {
	sawBefore := map[int]*Occurrence{}
	for idx, num := range data {
		sawBefore = addMapEntry(sawBefore, num, idx)
	}
	turn := len(data)
	lastSpoken := data[len(data) - 1]

	for turn < terminationTurn {
		if occurrence, _ := sawBefore[lastSpoken];
		occurrence.didSeeTwice() {
			//fmt.Printf("Turn %d - %d\n", turn+1, occurrence.diff())
			lastSpoken = occurrence.diff()
			sawBefore = addMapEntry(sawBefore, occurrence.diff(), turn)
		}else{
			//fmt.Printf("Turn %d - %d\n", turn + 1, 0)
			sawBefore = addMapEntry(sawBefore, 0, turn)
			lastSpoken = 0
		}
		turn += 1
	}
	return lastSpoken
}