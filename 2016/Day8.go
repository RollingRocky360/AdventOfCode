package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var screen [][]bool = [][]bool{}

const (
	screenWidth  int = 50
	screenHeight int = 6
)

var (
	rectRegex   = regexp.MustCompile(`^rect (\d+)x(\d+)$`)
	rotateRegex = regexp.MustCompile(`^rotate (row|column) [xy]=(\d+) by (\d+)$`)
)

func rect(width, height int) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			screen[h][w] = true
		}
	}
}

func rotateRow(row, reps int) {
	reps = reps % screenWidth
	anchor := len(screen[row]) - reps
	screen[row] = append(screen[row][anchor:], screen[row][:anchor]...)
}

func rotateCol(col, reps int) {
	reps = reps % screenHeight
	temp := make([]bool, screenHeight)

	for row := 0; row < screenHeight; row++ {
		temp[row] = screen[row][col]
	}

	anchor := len(temp) - reps
	temp = append(temp[anchor:], temp[:anchor]...)

	for row := 0; row < screenHeight; row++ {
		screen[row][col] = temp[row]
	}
}

func findOpAndExecute(instr string) {
	if matches := rectRegex.FindStringSubmatch(instr); len(matches) > 0 {
		width, _ := strconv.Atoi(matches[1])
		height, _ := strconv.Atoi(matches[2])
		rect(width, height)
	} else if matches := rotateRegex.FindStringSubmatch(instr); len(matches) > 0 {
		rowOrCol, _ := strconv.Atoi(matches[2])
		reps, _ := strconv.Atoi(matches[3])

		if matches[1] == "row" {
			rotateRow(rowOrCol, reps)
		} else {
			rotateCol(rowOrCol, reps)
		}
	}
}

func day8() {
	for cnt := 0; cnt < screenHeight; cnt++ {
		screen = append(screen, make([]bool, screenWidth))
	}

	inputFile, _ := os.Open("inputs/Day8.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		findOpAndExecute(scanner.Text())
	}

	count := 0
	for row := 0; row < screenHeight; row++ {
		for col := 0; col < screenWidth; col++ {
			if screen[row][col] {
				fmt.Print("#")
				count++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("\nnumber of pixels on =", count)

}
