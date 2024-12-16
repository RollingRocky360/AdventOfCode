package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	WALL  = byte('#')
	BOX   = byte('O')
	EMPTY = byte('.')
	ROBOT = byte('@')

	DBOXLEFT  = byte('[')
	DBOXRIGHT = byte(']')
	DBOX      = [2]byte{DBOXLEFT, DBOXRIGHT}
	DBOXINV   = [2]byte{DBOXRIGHT, DBOXLEFT}
)

type direction struct {
	horiz int
	vert  int
}

func Display(grid [][]byte, wait bool) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
	if !wait {
		return
	}
	time.Sleep(time.Millisecond * 200)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getRobotPos(grid [][]byte) (int, int) {
	for i, row := range grid {
		for j, c := range row {
			if c == ROBOT {
				return i, j
			}
		}
	}
	return -1, -1
}

func moveRobot(grid [][]byte, d direction, i, j int) (int, int) {
	ni, nj := i+d.vert, j+d.horiz

	if grid[ni][nj] == WALL {
		return i, j
	}

	if grid[ni][nj] == EMPTY {
		grid[ni][nj] = ROBOT
		grid[i][j] = EMPTY
		return ni, nj
	}

	for grid[ni][nj] == BOX {
		ni, nj = ni+d.vert, nj+d.horiz
	}

	if grid[ni][nj] == WALL {
		return i, j
	}

	grid[ni][nj] = BOX
	grid[i+d.vert][j+d.horiz] = ROBOT
	grid[i][j] = EMPTY

	return i + d.vert, j + d.horiz
}

func Part1() {
	grid, dirs := parseInput(false)
	i, j := getRobotPos(grid)

	for _, d := range dirs {
		i, j = moveRobot(grid, d, i, j)
	}

	totalGPS := 0
	for fromTop, row := range grid {
		for fromLeft, c := range row {
			if c == BOX {
				totalGPS += fromTop*100 + fromLeft
			}
		}
	}

	log.Println("part 1:", totalGPS)
}

func moveBoxHoriz(grid [][]byte, horiz, i, j int) {
	mul := 1
	if grid[i][j] == DBOXRIGHT {
		j--
	}
	if horiz > 0 {
		mul = 2
	}

	if grid[i][j+mul*horiz] == WALL {
		return
	}

	if grid[i][j+mul*horiz] == DBOXLEFT || grid[i][j+mul*horiz] == DBOXRIGHT {
		moveBoxHoriz(grid, horiz, i, j+mul*horiz)
	}

	if grid[i][j+mul*horiz] == EMPTY {
		grid[i][j+horiz] = DBOXLEFT
		if horiz > 0 {
			grid[i][j] = EMPTY
			grid[i][j+mul*horiz] = DBOXRIGHT
		} else {
			grid[i][j] = DBOXRIGHT
			grid[i][j+1] = EMPTY
		}
	}
}

func isMovable(grid [][]byte, vert, i, j int) bool {
	if grid[i][j] == DBOXRIGHT {
		j--
	}

	if grid[i+vert][j] == WALL || grid[i+vert][j+1] == WALL {
		return false
	}

	if grid[i+vert][j] == DBOXLEFT || grid[i+vert][j] == DBOXRIGHT {
		if !isMovable(grid, vert, i+vert, j) {
			return false
		}
	}
	if grid[i+vert][j+1] == DBOXLEFT || grid[i+vert][j+1] == DBOXRIGHT {
		if !isMovable(grid, vert, i+vert, j+1) {
			return false
		}
	}

	return true
}

func moveBoxVert(grid [][]byte, vert, i, j int) {
	if grid[i][j] == DBOXRIGHT {
		j--
	}

	if grid[i+vert][j] == WALL || grid[i+vert][j+1] == WALL {
		return
	}

	if grid[i+vert][j] == DBOXLEFT || grid[i+vert][j] == DBOXRIGHT {
		moveBoxVert(grid, vert, i+vert, j)
	}
	if grid[i+vert][j+1] == DBOXLEFT || grid[i+vert][j+1] == DBOXRIGHT {
		moveBoxVert(grid, vert, i+vert, j+1)
	}

	if grid[i+vert][j] == EMPTY && grid[i+vert][j+1] == EMPTY {
		grid[i+vert][j], grid[i+vert][j+1] = DBOXLEFT, DBOXRIGHT
		grid[i][j], grid[i][j+1] = EMPTY, EMPTY
	}
}

func moveWideBox(grid [][]byte, d direction, i, j int) {
	if d.vert == 0 {
		moveBoxHoriz(grid, d.horiz, i, j)
		return
	}

	if isMovable(grid, d.vert, i, j) {
		moveBoxVert(grid, d.vert, i, j)
	}
}

func moveRobotWide(grid [][]byte, d direction, i, j int) (int, int) {
	ni, nj := i+d.vert, j+d.horiz

	if grid[ni][nj] == WALL {
		return i, j
	}

	if grid[ni][nj] == DBOXLEFT || grid[ni][nj] == DBOXRIGHT {
		moveWideBox(grid, d, ni, nj)
	}

	if grid[ni][nj] == EMPTY {
		grid[ni][nj] = ROBOT
		grid[i][j] = EMPTY
		return ni, nj
	}

	return i, j
}

func Part2() {
	grid, dirs := parseInput(true)
	i, j := getRobotPos(grid)

	for _, d := range dirs {
		// Display(grid, true)
		i, j = moveRobotWide(grid, d, i, j)
	}

	totalGPS := 0
	for fromTop, row := range grid {
		for fromLeft, c := range row {
			if c == DBOXLEFT {
				totalGPS += fromTop*100 + fromLeft
			}
		}
	}

	log.Println("part 2:", totalGPS)
	// Display(grid, false)
}

func Run() {
	Part1()
	Part2()
}

func parseInput(doubleWide bool) ([][]byte, []direction) {
	input, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	grid := [][]byte{}
	for scanner.Scan() && len(scanner.Bytes()) != 0 {
		line := []byte{}
		for _, b := range scanner.Bytes() {
			if !doubleWide {
				line = append(line, b)
				continue
			}

			if b == ROBOT || b == EMPTY {
				line = append(line, b, EMPTY)
			}
			if b == BOX {
				line = append(line, DBOXLEFT, DBOXRIGHT)
			}
			if b == WALL {
				line = append(line, WALL, WALL)
			}
		}
		grid = append(grid, line)
	}

	dirs := []direction{}
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			var d direction
			switch c {
			case '^':
				d = direction{0, -1}
			case '<':
				d = direction{-1, 0}
			case 'v':
				d = direction{0, 1}
			case '>':
				d = direction{1, 0}
			}
			dirs = append(dirs, d)
		}
	}

	return grid, dirs
}
