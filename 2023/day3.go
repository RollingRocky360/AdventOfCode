package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var numRegex = regexp.MustCompile(`\d+`)

var (
	delx = []int{1, -1, 0, 0, 1, -1, 1, -1}
	dely = []int{0, 0, 1, -1, 1, -1, -1, 1}
)

func validPos(grid []string, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[1])
}

func parseNumberAndMark(grid []string, visited [][]bool, x, y int) (int, bool) {
	if !validPos(grid, x, y) || !unicode.IsDigit(rune(grid[x][y])) || visited[x][y] {
		return 0, false
	}
	left, right := y, y

	for left > 0 && unicode.IsDigit(rune(grid[x][left-1])) {
		left--
	}
	for right < len(grid[1])-1 && unicode.IsDigit(rune(grid[x][right+1])) {
		right++
	}

	for i := left; i <= right; i++ {
		visited[x][i] = true
	}

	num, _ := strconv.Atoi(string(grid[x][left : right+1]))
	return num, true
}

func getSumOfUnvisitedNum(grid []string, visited [][]bool, row, col int) int {
	sum := 0

	for i := 0; i < 8; i++ {
		x, y := delx[i], dely[i]
		if num, ok := parseNumberAndMark(grid, visited, row+x, col+y); ok {
			sum += num
		}
	}

	return sum
}

func getGearRatio(grid []string, visited [][]bool, row, col int) int {
	prod := 1
	cnt := 0

	for i := 0; i < 8; i++ {
		x, y := delx[i], dely[i]
		if num, ok := parseNumberAndMark(grid, visited, row+x, col+y); ok {
			prod *= num
			cnt++
		}
	}

	if cnt != 2 {
		return 0
	}
	return prod
}

func day3() {
	inputFile, err := os.Open("inputs/day3.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	res := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[1]))
	}

	for row, line := range grid {
		for col, ch := range line {
			if unicode.IsDigit(ch) || ch != '*' {
				continue
			}
			res += getGearRatio(grid, visited, row, col)
		}
	}

	fmt.Println(res)
}
