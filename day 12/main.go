package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	NorthStr   = "N"
	SouthStr   = "S"
	EastStr    = "E"
	WestStr    = "W"
	LeftStr    = "L"
	RightStr   = "R"
	ForwardStr = "F"
)

const Move = "NSEW"
const Rotate = "LR"

var TranslateDirection = map[string]IntVector{
	NorthStr: {x: 0, y: 1},
	SouthStr: {x: 0, y: -1},
	EastStr:  {x: 1, y: 0},
	WestStr:  {x: -1, y: 0},
}

var TranslateRotation = map[string]int{
	LeftStr:  -1,
	RightStr: 1,
}

var TranslateHeading = map[int]string{
	0:   NorthStr,
	90:  EastStr,
	180: SouthStr,
	270: WestStr,
}

func main() {
	input := readInput("./in.txt")
	part1(input)
	println("-------")
	part2(input)
}

func parseAction(line string) (string, int) {
	actionStr := line[0]
	magnitude, _ := strconv.Atoi(line[1:])
	return string(actionStr), magnitude
}

func makeMovePart1(actionStr string, heading int, position IntVector) (int, IntVector) {
	action, magnitude := parseAction(actionStr)
	switch {
	case strings.Contains(Move, action):
		position = position.add(TranslateDirection[action].mult(magnitude))
	case strings.Contains(Rotate, action):
		magnitude = magnitude % 360 // just in case there are values > 360
		heading = (heading + 360 + TranslateRotation[action]*magnitude) % 360
	case action == ForwardStr:
		position = position.add(TranslateDirection[TranslateHeading[heading]].mult(magnitude))
	}
	return heading, position
}

func part1(data []string) {
	heading := 90
	position := IntVector{x: 0, y: 0}
	for _, actionStr := range data {
		heading, position = makeMovePart1(actionStr, heading, position)
		if !(0 <= heading && heading < 360) {
			panic("Invalid heading")
		}
	}
	println("Part1:")
	fmt.Printf("Heading: %d, Position: %v\n", heading, position)
	println("Distance from start:", manhattanDist(position, IntVector{x: 0, y: 0}))
}

func makeMovePart2(actionStr string, ship, waypoint IntVector) (IntVector, IntVector) {
	action, magnitude := parseAction(actionStr)
	switch {
	case strings.Contains(Move, action):
		//	Move waypoint
		waypoint = waypoint.add(TranslateDirection[action].mult(magnitude))
	case strings.Contains(Rotate, action):
		// Rotate waypoint around ship
		magnitude = magnitude % 360 // just in case there are values > 360
		rotationDirection := action == LeftStr
		waypoint = waypoint.rotate(ship, float64(magnitude), rotationDirection)
	case action == ForwardStr:
		// move waypoint and ship to waypoint magnitude times
		// waypoint relative position stays the same
		waypointRelative := waypoint.add(ship.inverse())
		ship = ship.add(waypointRelative.mult(magnitude))
		waypoint = ship.add(waypointRelative)

	}
	return ship, waypoint
}

func part2(data []string) {
	ship := IntVector{x: 0, y: 0}
	waypoint := IntVector{x: 10, y: 1}
	for _, actionStr := range data {
		ship, waypoint = makeMovePart2(actionStr, ship, waypoint)
	}
	println("Part2:")
	fmt.Printf("Ship: %v, Waypoint, %v\n", ship, waypoint)
	println("Ship distance from start:", manhattanDist(ship, IntVector{x: 0, y: 0}))
}
