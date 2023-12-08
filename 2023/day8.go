package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	name  string
	left  string
	right string
}

func GCD(a, b int64) int64 {
	if b > a {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(a, b int64) int64 {
	return (a * b) / GCD(a, b)
}

func stepsToReachEnd(node int, endNodes map[int]bool, graph [][]int, dirs []int8) int64 {
	var steps int64

	for !endNodes[node] {
		for _, dir := range dirs {
			node = graph[node][dir]
			steps++
		}
	}

	return steps
}

func day8() {
	inputFile, err := os.Open("inputs/day8.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	dirs := []int8{}
	for _, dir := range scanner.Text() {
		if dir == 'L' {
			dirs = append(dirs, 0)
		} else {
			dirs = append(dirs, 1)
		}
	}

	scanner.Scan()
	nodeRegex := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)

	nodes := []Node{}
	startNodes := []int{}
	endNodes := map[int]bool{}

	nodeMap := map[string]int{}
	index := 0
	for scanner.Scan() {
		m := nodeRegex.FindStringSubmatch(scanner.Text())
		node, left, right := m[1], m[2], m[3]
		nodes = append(nodes, Node{node, left, right})
		nodeMap[node] = index

		switch node[2] {
		case 'A':
			startNodes = append(startNodes, index)
		case 'Z':
			endNodes[index] = true
		}

		index++
	}

	graph := [][]int{}
	for _, node := range nodes {
		left, right := nodeMap[node.left], nodeMap[node.right]
		graph = append(graph, []int{left, right})
	}

	steps := []int64{}
	for _, node := range startNodes {
		steps = append(steps, stepsToReachEnd(node, endNodes, graph, dirs))
	}

	var (
		lcm     int64 = 1
		product int64 = 1
	)
	for _, step := range steps {
		lcm = LCM(step, lcm)
		product *= step
	}

	fmt.Println(lcm)
}
