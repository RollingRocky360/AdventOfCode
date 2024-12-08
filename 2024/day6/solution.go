package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"time"
)

// -1 0 UP
// 0 1 RIGHT
// 1 0 DOWN
// 0 -1 LEFT

type direction struct {
	vert  int
	horiz int
}

func (d direction) encode() int {
	return 2*d.horiz - d.vert
}

const (
	MARK_VISIT = byte('X')
	OBSTACLE   = byte('#')
)

func getStartingPos(grid [][]byte) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == byte('^') {
				startI, startJ = i, j
				return i, j
			}
		}
	}

	for _, row := range grid {
		log.Println(string(row))
	}
	log.Fatal("could not find starting position")
	return 0, 0
}

func isWithinBounds(grid [][]byte, i, j int) bool {
	return i >= 0 &&
		i < len(grid) &&
		j >= 0 &&
		j < len(grid[0])
}

func turnRight90Deg(d direction) direction {
	return direction{
		vert:  d.horiz,
		horiz: d.vert * -1,
	}
}

func Part1() {
	grid := parseInput()
	i, j := getStartingPos(grid)
	dir := direction{vert: -1, horiz: 0}

	for {
		grid[i][j] = MARK_VISIT
		if !isWithinBounds(grid, i+dir.vert, j+dir.horiz) {
			break
		}
		if grid[i+dir.vert][j+dir.horiz] == OBSTACLE {
			dir = turnRight90Deg(dir)
		}
		i += dir.vert
		j += dir.horiz
	}

	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == MARK_VISIT {
				count++
			}
		}
	}

	log.Println("part 1:", count)
}

func getTools(grid [][]byte, dirInfo [][][]int) ([][]byte, [][][]int) {
	gridCopy := make([][]byte, len(grid))
	for i, row := range grid {
		gridCopy[i] = make([]byte, len(row))
		copy(gridCopy[i], row)
	}

	dirInfoCopy := make([][][]int, len(dirInfo))
	for i, row := range dirInfoCopy {
		dirInfoCopy[i] = make([][]int, len(dirInfo[0]))
		for j, cell := range row {
			dirInfoCopy[i][j] = make([]int, len(cell))
			copy(dirInfoCopy[i][j], cell)
		}
	}

	return gridCopy, dirInfoCopy
}

func canTrapGuard(g [][]byte, d [][][]int, i, j int, dir direction) bool {
	grid, dirInfo := getTools(g, d)

	if i+dir.vert == startI && j+dir.horiz == startJ {
		return false
	}

	grid[i+dir.vert][j+dir.horiz] = OBSTACLE

	i, j = startI, startJ
	dir = direction{-1, 0}
	for {
		// Display(grid)

		grid[i][j] = MARK_VISIT
		dirInfo[i][j] = append(dirInfo[i][j], dir.encode())

		nextI, nextJ := i+dir.vert, j+dir.horiz

		if !isWithinBounds(grid, nextI, nextJ) {
			return false
		}

		for isWithinBounds(grid, nextI, nextJ) && grid[nextI][nextJ] == OBSTACLE {
			dir = turnRight90Deg(dir)
			nextI, nextJ = i+dir.vert, j+dir.horiz
		}

		i += dir.vert
		j += dir.horiz

		if grid[i][j] == MARK_VISIT && slices.Contains(dirInfo[i][j], dir.encode()) {
			return true
		}
	}
}

func Display(grid [][]byte) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for _, row := range grid {
		fmt.Println(string(row))
	}
	time.Sleep(time.Millisecond * 20)
}

var (
	startI int
	startJ int
)

func Part2() {
	grid := parseInput()
	i, j := getStartingPos(grid)
	dirInfo := make([][][]int, len(grid))
	for x := range grid {
		dirInfo[x] = make([][]int, len(grid[0]))
	}
	dir := direction{vert: -1, horiz: 0}

	traps := map[[2]int]bool{}

	for {
		grid[i][j] = MARK_VISIT
		dirInfo[i][j] = append(dirInfo[i][j], dir.encode())
		nextI, nextJ := i+dir.vert, j+dir.horiz

		for isWithinBounds(grid, nextI, nextJ) && grid[nextI][nextJ] == OBSTACLE {
			dir = turnRight90Deg(dir)
			nextI, nextJ = i+dir.vert, j+dir.horiz
		}

		if !isWithinBounds(grid, nextI, nextJ) {
			break
		}

		if canTrapGuard(grid, dirInfo, i, j, dir) {
			traps[[2]int{nextI, nextJ}] = true
		}

		i += dir.vert
		j += dir.horiz
	}

	log.Println("part 2:", len(traps))
}

func Run() {
	Part1()
	Part2()
}

func parseInput() [][]byte {
	input, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		scanned := scanner.Bytes()
		line := make([]byte, len(scanned))
		copy(line, scanned)
		grid = append(grid, line)
	}

	return grid
}
