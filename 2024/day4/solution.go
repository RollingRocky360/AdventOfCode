package day4

import (
	"bufio"
	"log"
	"os"
)

const (
	DIR_RIGHT = 1
	DIR_LEFT  = -1
	DIR_DOWN  = 1
	DIR_UP    = -1
	DIR_NOP   = 0
)

var (
	DIRS = [8][2]int{
		{DIR_RIGHT, DIR_NOP},
		{DIR_RIGHT, DIR_DOWN},
		{DIR_NOP, DIR_DOWN},
		{DIR_LEFT, DIR_DOWN},
		{DIR_LEFT, DIR_NOP},
		{DIR_LEFT, DIR_UP},
		{DIR_NOP, DIR_UP},
		{DIR_RIGHT, DIR_UP},
	}
)

var (
	XMAS = []byte("XMAS")
)

func explore(grid [][]byte, x, y, rightStep, downStep, pos int) int {
	if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0 {
		return 0
	}

	if grid[x][y] != XMAS[pos] {
		return 0
	}

	if pos == 3 {
		return 1
	}

	return explore(grid, x+downStep, y+rightStep, rightStep, downStep, pos+1)
}

func exploreXMAS(grid [][]byte, x, y int) int {
	count := 0
	for _, dir := range DIRS {
		count += explore(grid, x, y, dir[0], dir[1], 0)
	}
	return count
}

func exploreCrossMAS(grid [][]byte, x, y int) int {
	count := 0

	for i := 1; i < len(DIRS); i += 2 {
		dir := DIRS[i]
		right, down := dir[0], dir[1]

		if explore(grid, x, y, right, down, 1) == 1 &&
			(explore(grid, x, y+2*right, -right, down, 1) > 0 || explore(grid, x+2*down, y, right, -down, 1) > 0) {
			count++
		}
	}

	return count
}

func Part1() {
	grid := parseInput()

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count += exploreXMAS(grid, i, j)
		}
	}

	log.Println("part 1:", count)
}

func Part2() {
	grid := parseInput()

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count += exploreCrossMAS(grid, i, j)
		}
	}

	log.Println("part 2:", count/2)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() [][]byte {
	input, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	wordSearch := [][]byte{}

	for scanner.Scan() {
		scanned := scanner.Bytes()
		line := make([]byte, len(scanned))
		copy(line, scanned)
		wordSearch = append(
			wordSearch, line,
		)
	}

	return wordSearch
}
