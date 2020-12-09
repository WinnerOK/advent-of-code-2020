package main

import "fmt"

func part1(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 2020 {
				fmt.Printf("%d\n", data[i]*data[j])
			}
		}
	}
}

func part2(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == 2020 {
					fmt.Printf("%d\n", data[i]*data[j]*data[k])
				}
			}
		}
	}
}

func main() {
	input := readInput("./in.txt")
	intInput := stringSliceToIntSlice(input)
	part1(intInput)
	part2(intInput)
}
