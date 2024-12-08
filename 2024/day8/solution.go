package day8

import (
	"bufio"
	"log"
	"os"
)

type pos struct {
	i int
	j int
}

const EMPTY = byte('.')

func isInBounds(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) &&
		j >= 0 && j < len(grid[0])
}

func getAntiNodesWithDiff2(grid [][]byte, towerPositions []pos) []pos {

	antinodes := []pos{}
	for i, first := range towerPositions[:len(towerPositions)-1] {
		for _, second := range towerPositions[i+1:] {
			antiI1 := 2*first.i - second.i
			antiJ1 := 2*first.j - second.j

			if isInBounds(grid, antiI1, antiJ1) {
				antinodes = append(antinodes, pos{antiI1, antiJ1})
			}
			antiI2 := 2*second.i - first.i
			antiJ2 := 2*second.j - first.j

			if isInBounds(grid, antiI2, antiJ2) {
				antinodes = append(antinodes, pos{antiI2, antiJ2})
			}
		}
	}

	return antinodes
}

func Part1() {
	grid := parseInput()
	towers := map[byte][]pos{}

	for i, row := range grid {
		for j, spot := range row {
			if spot != EMPTY {
				towers[spot] = append(towers[spot], pos{i, j})
			}
		}
	}

	uniqueAntinodes := map[pos]bool{}

	for _, t := range towers {
		for _, antinode := range getAntiNodesWithDiff2(grid, t) {
			uniqueAntinodes[antinode] = true
		}
	}

	log.Println("part 1:", len(uniqueAntinodes))
}

func getAntiNodes(grid [][]byte, towerPositions []pos) []pos {

	antinodes := []pos{}
	for i, first := range towerPositions[:len(towerPositions)-1] {
		for _, second := range towerPositions[i+1:] {
			diffI := first.i - second.i
			diffJ := first.j - second.j

			d := 0
			for {
				antiI1 := first.i + diffI*d
				antiJ1 := first.j + diffJ*d

				hit := false

				if isInBounds(grid, antiI1, antiJ1) {
					hit = true
					antinodes = append(antinodes, pos{antiI1, antiJ1})
				}

				antiI2 := first.i - diffI*d
				antiJ2 := first.j - diffJ*d

				if isInBounds(grid, antiI2, antiJ2) {
					hit = true
					antinodes = append(antinodes, pos{antiI2, antiJ2})
				}

				d++
				if !hit {
					break
				}
			}
		}
	}

	return antinodes
}

func Part2() {
	grid := parseInput()
	towers := map[byte][]pos{}

	for i, row := range grid {
		for j, spot := range row {
			if spot != EMPTY {
				towers[spot] = append(towers[spot], pos{i, j})
			}
		}
	}

	uniqueAntinodes := map[pos]bool{}

	for _, t := range towers {
		for _, antinode := range getAntiNodes(grid, t) {
			uniqueAntinodes[antinode] = true
		}
	}

	for _, t := range towers {
		for _, antinode := range getAntiNodesWithDiff2(grid, t) {
			uniqueAntinodes[antinode] = true
		}
	}

	log.Println("part 2:", len(uniqueAntinodes))
}

func Run() {
	Part1()
	Part2()
}

func parseInput() [][]byte {
	input, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	grid := [][]byte{}
	for scanner.Scan() {
		scanned := scanner.Bytes()
		line := make([]byte, len(scanned))
		copy(line, scanned)
		grid = append(grid, line)
	}

	return grid
}
