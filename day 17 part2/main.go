package main

const (
	ActiveCube   = '#'
	InactiveCube = '.'
)

// xyz
type Coordinate struct {
	x, y, z, w int
}
type cubeMap map[Coordinate]bool

func copyCubeMap(m cubeMap) cubeMap {
	newMap := cubeMap{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func parseMap(map2d []string) cubeMap {
	result := cubeMap{}
	for y, line := range map2d {
		for x, cube := range line {
			if cube == ActiveCube {
				result[Coordinate{
					x: x,
					y: y,
					z: 0,
					w: 0,
				}] = true
			}
		}
	}
	return result
}

func main() {
	input := readInput("./in.txt")
	initialMap := parseMap(input)
	part2(copyCubeMap(initialMap))
}

func generateNeighbours(coordinate Coordinate) []Coordinate {
	result := []Coordinate{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if !(dx == 0 && dy == 0 && dz == 0 && dw == 0) {
						result = append(result, Coordinate{
							x: coordinate.x + dx,
							y: coordinate.y + dy,
							z: coordinate.z + dz,
							w: coordinate.w + dw,
						})
					}
				}
			}
		}
	}
	return result
}

func countActiveNeighbours(m cubeMap, coordinate Coordinate) int {
	activeNeighbours := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if !(dx == 0 && dy == 0 && dz == 0 && dw == 0) {
						if val, present := m[Coordinate{
							x: coordinate.x + dx,
							y: coordinate.y + dy,
							z: coordinate.z + dz,
							w: coordinate.w + dw,
						}]; present && val {
							activeNeighbours += 1
						}
					}
				}
			}
		}
	}
	return activeNeighbours
}

func part2(m cubeMap) {
	for cycle := 1; cycle <= 6; cycle++ {
		nextMap := copyCubeMap(m)
		for coordinate, active := range m {
			// for current active node
			if active {
				activeNeighbours := countActiveNeighbours(m, coordinate)
				if !(activeNeighbours == 2 || activeNeighbours == 3) {
					// if node was active and there is no more 2 or 3 active neighbours - die
					delete(nextMap, coordinate)
				}
			}
			for _, neighbourCoord := range generateNeighbours(coordinate) {
				if val, ok := m[neighbourCoord]; !ok || !val {
					//	consider inactive neighbour only. Active will be considered in outer loop
					activeNeighbours := countActiveNeighbours(m, neighbourCoord)
					if activeNeighbours == 3 {
						nextMap[neighbourCoord] = true
					}
				}
			}
		}
		m = nextMap
	}
	println("Part 2:", len(m))
}
