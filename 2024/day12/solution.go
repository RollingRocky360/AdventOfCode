package day12

import (
	"bufio"
	"log"
	"os"
)

type Region struct {
	perimeter int
	area      int
	corners   int
}

type direction struct {
	horiz int
	vert  int
}

type state struct {
	plant        byte
	exploreCount int
}

var DIRS = [4]direction{
	{-1, 0}, // LEFT
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
}

var CYCLIC_DIRS = [5]direction{
	{-1, 0}, // LEFT
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
}

var registry = map[state]Region{}
var visited = map[[2]int]bool{}

func isInBounds(grid [][]byte, i, j int) bool {
	if i < 0 || i >= len(grid) {
		return false
	}
	if j < 0 || j >= len(grid[0]) {
		return false
	}
	return true
}

func explore(grid [][]byte, s state, i, j int) {
	if visited[[2]int{i, j}] {
		return
	}

	visited[[2]int{i, j}] = true

	r := registry[s]
	r.area++
	registry[s] = r

	for _, dir := range DIRS {

		nextI, nextJ := i+dir.vert, j+dir.horiz
		if !isInBounds(grid, nextI, nextJ) || grid[nextI][nextJ] != s.plant {
			r := registry[s]
			r.perimeter++
			registry[s] = r
			continue
		}

		explore(grid, s, nextI, nextJ)
	}
}

func Part1() {
	grid := parseInput()

	exploreCount := 0
	for i, row := range grid {
		for j, plant := range row {
			explore(grid, state{plant, exploreCount}, i, j)
			exploreCount++
		}
	}

	cost := 0
	for _, r := range registry {
		cost += r.area * r.perimeter
	}

	log.Println("part 1:", cost)
}

func inwardCorner(grid [][]byte, s state, i, j int) int {
	count := 0
	for x := 0; x < len(CYCLIC_DIRS)-1; x++ {
		d1, d2 := CYCLIC_DIRS[x], CYCLIC_DIRS[x+1]

		if !isInBounds(grid, i+d1.vert, j+d1.horiz) || !isInBounds(grid, i+d2.vert, j+d2.horiz) {
			continue
		}

		diagVert, diagHoriz := d1.vert+d2.vert, d1.horiz+d2.horiz
		if grid[i+diagVert][j+diagHoriz] == s.plant {
			continue
		}

		if grid[i+d1.vert][j+d1.horiz] == grid[i+d2.vert][j+d2.horiz] && grid[i+d2.vert][j+d2.horiz] == s.plant {
			count++
		}
	}

	return count
}

func outwardCorner(grid [][]byte, s state, i, j int) int {
	count := 0
	for x := 0; x < len(CYCLIC_DIRS)-1; x++ {
		d1, d2 := CYCLIC_DIRS[x], CYCLIC_DIRS[x+1]

		n1 := !isInBounds(grid, i+d1.vert, j+d1.horiz) || grid[i+d1.vert][j+d1.horiz] != s.plant
		n2 := !isInBounds(grid, i+d2.vert, j+d2.horiz) || grid[i+d2.vert][j+d2.horiz] != s.plant

		if n1 && n2 {
			count++
		}
	}

	return count
}

func exploreSides(grid [][]byte, s state, i, j int) {
	if visited[[2]int{i, j}] {
		return
	}

	visited[[2]int{i, j}] = true

	r := registry[s]
	r.area++
	registry[s] = r

	for _, dir := range DIRS {
		nextI, nextJ := i+dir.vert, j+dir.horiz

		if !isInBounds(grid, nextI, nextJ) || grid[nextI][nextJ] != s.plant {
			continue
		}

		exploreSides(grid, s, nextI, nextJ)
	}

	r = registry[s]
	outwards := outwardCorner(grid, s, i, j)
	inwards := inwardCorner(grid, s, i, j)
	r.corners += outwards + inwards
	registry[s] = r

}

func Part2() {
	grid := parseInput()

	registry = map[state]Region{}
	visited = map[[2]int]bool{}

	exploreCount := 0
	for i, row := range grid {
		for j, plant := range row {
			exploreSides(grid, state{plant, exploreCount}, i, j)
			exploreCount++
		}
	}

	cost := 0
	for _, r := range registry {
		cost += r.area * r.corners
	}

	log.Println("part 2:", cost)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() [][]byte {
	input, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	grid := [][]byte{}

	for scanner.Scan() {
		row := scanner.Bytes()
		rowCopy := make([]byte, len(row))
		copy(rowCopy, row)
		grid = append(grid, rowCopy)
	}

	return grid
}
