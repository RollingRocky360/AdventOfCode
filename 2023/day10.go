package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Point struct {
	x int
	y int
}

const (
	pipe     = 1
	enclosed = 2
)

// expand grid and fill gaps with either - or | appropriately
// followed by flood-fill
func day10_sol1() {
	island := parseIsland()

	start, ok := fetchStartingPoint(island)
	if !ok {
		panic("No starting point found")
	}

	pos1 := Point{start.x + 1, start.y}
	prev1 := start
	stepCount := 1

	islandMap := make([][]int, len(island)*2)
	for i := 0; i < len(islandMap); i++ {
		islandMap[i] = make([]int, len(island[0])*2)
	}

	islandMap[start.x*2][start.y*2] = pipe
	islandMap[(start.x*2)+1][start.y*2] = pipe
	islandMap[pos1.x*2+1][pos1.y*2] = pipe
	islandMap[(pos1.x*2)+1][pos1.y*2] = pipe

	visited := make([][]bool, len(islandMap))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(islandMap[0]))
	}
	visited[start.x*2][start.y*2] = true
	visited[pos1.x*2][pos1.y*2] = true
	visited[start.x*2+1][start.y*2] = true
	visited[pos1.x*2+1][pos1.y*2] = true

	for pos1 != start {
		temp1 := pos1
		delX, delY := move(island, pos1, prev1)
		pos1.x += delX
		pos1.y += delY
		prev1 = temp1
		stepCount++

		islandMap[pos1.x*2][pos1.y*2] = pipe
		islandMap[pos1.x*2-delX][pos1.y*2-delY] = pipe
		visited[pos1.x*2][pos1.y*2] = true
		visited[pos1.x*2-delX][pos1.y*2-delY] = true
	}

	fmt.Println(stepCount / 2)

	bfs(islandMap, visited, Point{0, 0})
	bfs(islandMap, visited, Point{len(islandMap) - 1, 0})
	bfs(islandMap, visited, Point{0, len(islandMap[0]) - 1})
	bfs(islandMap, visited, Point{len(islandMap) - 1, len(islandMap[0]) - 1})

	totalEnclosed := 0
	for i := 0; i < len(islandMap); i += 2 {
		for j := 0; j < len(islandMap[0]); j += 2 {
			if !visited[i][j] {
				totalEnclosed++
			}
		}
	}

	fmt.Println(totalEnclosed)
}

// replace L-*7 and F-*J shaped walls with |
// and perform row-wise parity scan (pipe wall hit)
func day10_sol2() {
	island := parseIsland()

	start, ok := fetchStartingPoint(island)
	if !ok {
		panic("No starting point found")
	}

	pos1 := Point{start.x + 1, start.y}
	prev1 := start
	stepCount := 1

	wallMap := make([][]rune, len(island))
	for i := 0; i < len(wallMap); i++ {
		wallMap[i] = make([]rune, len(island[0]))
		for j := 0; j < len(wallMap[i]); j++ {
			wallMap[i][j] = '.'
		}
	}

	wallMap[pos1.x][pos1.y] = island[pos1.x][pos1.y]

	for pos1 != start {
		temp1 := pos1
		delX, delY := move(island, pos1, prev1)
		pos1.x += delX
		pos1.y += delY
		prev1 = temp1
		stepCount++

		wallMap[pos1.x][pos1.y] = island[pos1.x][pos1.y]
	}

	wallMap[start.x][start.y] = 'F'

	L7WallRegex := regexp.MustCompile(`(L-*7)|(F-*J)`)
	totalEnclosed := 0
	for _, line := range wallMap {
		morphed := L7WallRegex.ReplaceAllString(string(line), "|")
		wallHitCount := 0
		for _, tile := range morphed {
			if tile == '|' {
				wallHitCount++
			}
			if strings.ContainsRune("F|J-7LS", tile) {
				continue
			}
			if wallHitCount%2 == 1 {
				totalEnclosed++
			}
		}
	}

	fmt.Println(totalEnclosed)
}

func bfs(islandMap [][]int, visited [][]bool, p Point) {
	queue := []Point{p}
	dirx, diry := []int{0, 0, 1, -1}, []int{1, -1, 0, 0}

	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]

		x, y := point.x, point.y
		for i := 0; i < 4; i++ {
			newX, newY := x+dirx[i], y+diry[i]
			if newX < 0 || newX >= len(islandMap) || newY < 0 || newY >= len(islandMap[0]) || visited[newX][newY] {
				continue
			}
			queue = append(queue, Point{newX, newY})
			visited[newX][newY] = true
		}
	}
}

func move(island [][]rune, p Point, prev Point) (int, int) {
	var delX, delY int

	switch island[p.x][p.y] {
	case '|':
		if p.x-1 == prev.x {
			delX = 1
		} else {
			delX = -1
		}
	case '-':
		if p.y-1 == prev.y {
			delY = 1
		} else {
			delY = -1
		}
	case 'L':
		if p.y+1 == prev.y {
			delX = -1
		} else {
			delY = 1
		}
	case 'J':
		if p.y-1 == prev.y {
			delX = -1
		} else {
			delY = -1
		}
	case '7':
		if p.x+1 == prev.x {
			delY = -1
		} else {
			delX = 1
		}
	case 'F':
		if p.y+1 == prev.y {
			delX = 1
		} else {
			delY = 1
		}
	}

	return delX, delY
}

func parseIsland() [][]rune {
	inputFile, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err.Error())
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	island := [][]rune{}
	for scanner.Scan() {
		line := []rune{}
		for _, tile := range scanner.Text() {
			line = append(line, tile)
		}
		island = append(island, line)
	}

	return island
}

func fetchStartingPoint(island [][]rune) (Point, bool) {
	for x, line := range island {
		for y, pipe := range line {
			if pipe == 'S' {
				return Point{x, y}, true
			}
		}
	}

	return Point{}, false
}
