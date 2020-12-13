package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./in.txt")
	arrivalTime, _ := strconv.Atoi(input[0])
	busData := strings.Split(input[1], ",")
	busIDs := make([]int, len(busData))
	for idx, busEntry := range busData {
		if id, err := strconv.Atoi(busEntry); err == nil {
			busIDs[idx] = id
		} else {
			busIDs[idx] = -1
		}
	}
	part1(arrivalTime, busIDs)
	part2(busIDs)
}

func part1(arrivalTime int, busIDs []int) {
	minDepartureDelay := math.MaxInt32
	bestBus := -1
	for _, busEntry := range busIDs {
		if busEntry != -1 {
			if arrivalTime%busEntry == 0 {
				bestBus = busEntry
				minDepartureDelay = 0
				break
			} else {
				currentDelay := busEntry - (arrivalTime % busEntry)
				if currentDelay < minDepartureDelay {
					bestBus = busEntry
					minDepartureDelay = currentDelay
				}
			}
		}
	}

	if bestBus == -1 {
		panic("No bus found")
	}

	println("Part 1:", bestBus*minDepartureDelay)
}

var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	// Implementation of chinese remainder theorem from
	// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func part2(busIDs []int){
	n := []*big.Int{}
	a := []*big.Int{}
	// (x + late) mod busID == 0
	// x mod busID == busID - (idx % busID)
	for late, busID := range busIDs {
		if busID != -1 {
			a = append(a, big.NewInt(int64(busID - late%busID)))
			n = append(n, big.NewInt(int64(busID)))
		}
	}
	answer, err := crt(a,n)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:",  answer)
}

