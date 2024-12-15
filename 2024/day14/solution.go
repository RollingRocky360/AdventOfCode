package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	LEN_X = 101
	LEN_Y = 103
)

type Robot struct {
	px int
	py int
	vx int
	vy int
}

func (r *Robot) updatePos() {
	r.px = ((r.px+r.vx)%LEN_X + LEN_X) % LEN_X
	r.py = ((r.py+r.vy)%LEN_Y + LEN_Y) % LEN_Y
}

func Display(robots []*Robot) {
	grid := make([][]bool, LEN_Y)
	for i := range grid {
		grid[i] = make([]bool, LEN_X)
	}

	for _, r := range robots {
		grid[r.py][r.px] = true
	}

	for _, row := range grid {
		for _, point := range row {
			if point {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	time.Sleep(time.Millisecond * 100)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getQuadrant(x, y int) int {
	if x < LEN_X/2 && y < LEN_Y/2 {
		return 0
	}
	if x > LEN_X/2 && y < LEN_Y/2 {
		return 1
	}
	if x < LEN_X/2 && y > LEN_Y/2 {
		return 2
	}
	if x > LEN_X/2 && y > LEN_Y/2 {
		return 3
	}
	return 4
}

func Part1() {
	robots := parseInput()

	for _, r := range robots {
		for range 100 {
			r.updatePos()
		}
	}

	counts := [5]int{}
	for _, r := range robots {
		q := getQuadrant(r.px, r.py)
		counts[q]++
	}

	prod := 1
	for _, count := range counts[:4] {
		prod *= count
	}

	log.Println("part 1:", prod)
}

func checkFor3x3(robots []*Robot) bool {
	grid := make([][]int, LEN_Y)
	for i := range grid {
		grid[i] = make([]int, LEN_X)
	}

	for _, r := range robots {
		grid[r.py][r.px] = 1
	}

	for y := range LEN_Y - 3 {
		for x := range LEN_X - 3 {
			count := 0
			for dy := range 3 {
				for dx := range 3 {
					count += grid[y+dy][x+dx]
				}
			}

			if count == 9 {
				return true
			}
		}
	}

	return false
}

func Part2() {
	robots := parseInput()

	iter := 0
	for {
		for _, r := range robots {
			r.updatePos()
		}
		iter++
		// Display(robots)

		// Christmas tree appears if a 3x3 block is fully filled
		// observation by https://github.com/milindmadhukar
		if checkFor3x3(robots) {
			break
		}
	}

	log.Println("part 2:", iter)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() []*Robot {
	input, err := os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	robots := []*Robot{}
	for scanner.Scan() {
		r := Robot{}
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)
		robots = append(robots, &r)
	}

	return robots
}
