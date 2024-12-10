package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction struct {
	horiz int
	vert  int
}

var DIRS = [4]direction{
	{-1, 0}, // LEFT
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
}

func isInBounds(grid [][]int, i, j int) bool {
	if i < 0 || i >= len(grid) {
		return false
	}
	if j < 0 || j >= len(grid[0]) {
		return false
	}
	return true
}

func getScoreNoRepeat(grid [][]int, visited map[[2]int]bool, i, j, prevHeight int) int {
	if !isInBounds(grid, i, j) {
		return 0
	}

	if grid[i][j] != prevHeight+1 {
		return 0
	}

	if grid[i][j] == 9 && !visited[[2]int{i, j}] {
		visited[[2]int{i, j}] = true
		return 1
	}

	score := 0
	for _, dir := range DIRS {
		score += getScoreNoRepeat(grid, visited, i+dir.vert, j+dir.horiz, grid[i][j])
	}

	return score
}

func Part1() {
	grid := parseInput()

	totalScore := 0

	for i, row := range grid {
		for j, height := range row {
			if height == 0 {
				visited := map[[2]int]bool{}
				totalScore += getScoreNoRepeat(grid, visited, i, j, -1)
			}
		}
	}

	log.Println("part 1:", totalScore)
}

func getScore(grid [][]int, i, j, prevHeight int) int {
	if !isInBounds(grid, i, j) {
		return 0
	}

	if grid[i][j] != prevHeight+1 {
		return 0
	}

	if grid[i][j] == 9 {
		return 1
	}

	score := 0
	for _, dir := range DIRS {
		score += getScore(grid, i+dir.vert, j+dir.horiz, grid[i][j])
	}

	return score
}

func Part2() {
	grid := parseInput()

	totalScore := 0

	for i, row := range grid {
		for j, height := range row {
			if height == 0 {
				totalScore += getScore(grid, i, j, -1)
			}
		}
	}

	log.Println("part 2:", totalScore)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() [][]int {
	input, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	grid := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, heightText := range strings.Split(scanner.Text(), "") {
			height, err := strconv.Atoi(heightText)
			if err != nil {
				height = -1
			}
			row = append(row, height)
		}
		grid = append(grid, row)
	}

	return grid
}
