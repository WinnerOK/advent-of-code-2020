package main

import (
	"regexp"
	"strconv"
	"strings"
)

var bagsInsideDescription = regexp.MustCompile(`((?P<count>\d+) (?P<bag>[a-z ]+) bags?)[,\.]`)

type edge struct {
	src, dest string
	weight    int
}

func parseLine(line string) []edge {
	lineSplit := strings.Split(line, " bags contain")
	base := lineSplit[0]
	insideData := lineSplit[1]
	bagsInside := bagsInsideDescription.FindAllStringSubmatch(insideData, -1)
	result := []edge{}
	//2 - count, 3 - type
	for _, parsedEdge := range bagsInside {
		weight, _ := strconv.Atoi(parsedEdge[2])
		result = append(result, edge{
			src:    base,
			dest:   parsedEdge[3],
			weight: weight,
		})
	}
	return result
}

var used = make(map[string]bool)

func dfs(graph map[string]map[string]int, v string){
	if _, ok := used[v]; ok{
		return
	}else{
		used[v] = true
		for neighbour, _ := range graph[v]{
			dfs(graph, neighbour)
		}
	}
}

func dfs2(graph map[string]map[string]int, v string) int {
	countInside := 0
	for neighbour, amount := range graph[v] {
		countInside += amount * dfs2(graph, neighbour)
	}
	return 1 + countInside
}

func solve(data []string) {
	directGraph := make(map[string]map[string]int)
	inverseGraph := make(map[string]map[string]int)

	for _, line := range data {
		edges := parseLine(line)
		for _, edge := range edges {
			if _, ok := directGraph[edge.src]; !ok {
				directGraph[edge.src] = make(map[string]int)
			}
			if _, ok := inverseGraph[edge.dest]; !ok {
				inverseGraph[edge.dest] = make(map[string]int)
			}
			directGraph[edge.src][edge.dest] = edge.weight
			inverseGraph[edge.dest][edge.src] = edge.weight
		}
	}

	dfs(inverseGraph, "shiny gold")
	println("Part 1:", len(used) - 1)
	println("Part 2:", dfs2(directGraph, "shiny gold") - 1)

}

func main() {
	input := readInput("./in.txt")
	solve(input)
}
