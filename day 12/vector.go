package main

import "math"

type IntVector struct {
	x, y int
}

func (v IntVector) add(other IntVector) IntVector {
	return IntVector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v IntVector) mult(size int) IntVector {
	return IntVector{
		x: v.x * size,
		y: v.y * size,
	}
}

func (v IntVector) inverse() IntVector {
	return IntVector{
		x: -v.x,
		y: -v.y,
	}
}

func (v IntVector) rotate(origin IntVector, deg float64, counterClockWise bool) IntVector {
	relative := v.add(origin.inverse())
	if !counterClockWise {
		deg = - deg
	}
	deg = deg * math.Pi / 180
	newX := float64(relative.x)*math.Cos(deg) - float64(relative.y)*math.Sin(deg)
	newY := float64(relative.x)*math.Sin(deg) + float64(relative.y)*math.Cos(deg)

	newRelative := IntVector{
		x: int(math.Round(newX)),
		y: int(math.Round(newY)),
	}

	return newRelative.add(origin)
}

func manhattanDist(v1, v2 IntVector) int {
	return abs(v1.x-v2.x) + abs(v1.y-v2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
