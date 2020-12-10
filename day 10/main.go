package main

const (
	maxdiff = 3
)

func main() {
	input := readInput("./in.txt")
	intInput := stringSliceToIntSlice(input)
	part1(intInput)
	part2(intInput)
}

func part1(data []int) {
	adapters := map[int]int{}
	maxJoltage := 0
	for _, adapter := range data {
		adapters[adapter] = adapter
		if adapter > maxJoltage {
			maxJoltage = adapter
		}
	}
	myAdapter := maxJoltage + maxdiff
	adapters[myAdapter] = myAdapter
	currentJoltage := 0
	diffs := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}
	_, diffmap := getNextAdapterDiff(adapters, currentJoltage, diffs)
	println(diffmap[1], diffmap[2], diffmap[3], " Part 1: ", diffmap[1]*diffmap[3])
}

func getNextAdapterDiff(adapters map[int]int, currentJoltage int, currentDifs map[int]int) (int, map[int]int) {
	for diff := 1; diff <= maxdiff; diff++ {
		if adapterJoltage, ok := adapters[currentJoltage+diff]; ok {
			currentDifs[diff] += 1
			insideDiff, insideMap := getNextAdapterDiff(adapters, adapterJoltage, currentDifs)
			return diff + insideDiff, insideMap
		}
	}
	return 0, currentDifs
}

func part2(data []int) {
	adapters := map[int]int{}
	maxJoltage := 0
	for _, adapter := range data {
		adapters[adapter] = adapter
		if adapter > maxJoltage {
			maxJoltage = adapter
		}
	}
	myAdapter := maxJoltage + maxdiff
	adapters[myAdapter] = myAdapter
	adapters[0] = 0

	// Build a graph
	currentJoltage := 0
	graph := make(map[int]map[int]bool)
	for currentJoltage != myAdapter {
		if _, currentPresent := adapters[currentJoltage]; currentPresent {
			for diff := 1; diff <= maxdiff; diff++ {
				if _, adapterPresent := adapters[currentJoltage+diff]; adapterPresent {
					if _, mapPresent := graph[currentJoltage]; !mapPresent {
						graph[currentJoltage] = make(map[int]bool)
					}
					graph[currentJoltage][currentJoltage+diff] = true
					//println(currentJoltage, " -> ", currentJoltage+diff)
				}
			}
		}
		currentJoltage += 1
	}
	println("Part 2: ", countPaths(graph, myAdapter, 0))
}

var dist = map[int]int{}
var found = map[int]bool{}

func count(graph map[int]map[int]bool, v int) int {
	if found[v] {
		return dist[v]
	} else {
		sum := 0
		found[v] = true
		for neighbour, _ := range graph[v] {
			sum += count(graph, neighbour)
		}
		dist[v] = sum
		return sum
	}
}

func countPaths(g map[int]map[int]bool, s int, t int) int {
	// https://neerc.ifmo.ru/wiki/index.php?title=%D0%97%D0%B0%D0%B4%D0%B0%D1%87%D0%B0_%D0%BE_%D1%87%D0%B8%D1%81%D0%BB%D0%B5_%D0%BF%D1%83%D1%82%D0%B5%D0%B9_%D0%B2_%D0%B0%D1%86%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%BE%D0%BC_%D0%B3%D1%80%D0%B0%D1%84%D0%B5
	found[s] = true
	dist[s] = 1
	return count(g, t)
}
