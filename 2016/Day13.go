package main

import "fmt"

type position struct {
	x int
	y int
}

const favNumber = 1352

var destination = position{31, 39}

func isOpenSpace(x, y int) bool {
	sum := x*x + 3*x + 2*x*y + y + y*y + favNumber
	count := 0

	for sum > 0 {
		count += sum & 1
		sum >>= 1
	}

	return count%2 == 0
}

func day13() {
	levelEndMarker := position{-1, -1}
	queue := []position{{1, 1}, levelEndMarker}
	visited := map[position]bool{}
	stepCount := 0

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		x, y := pos.x, pos.y

		if pos == levelEndMarker {
			queue = append(queue, levelEndMarker)
			stepCount++
			if stepCount == 51 {
				break
			}
		}

		if x < 0 || y < 0 || !isOpenSpace(x, y) || visited[pos] {
			continue
		}

		visited[pos] = true

		queue = append(
			queue,
			position{x - 1, y},
			position{x + 1, y},
			position{x, y - 1},
			position{x, y + 1},
		)
	}

	fmt.Println(len(visited))
}
