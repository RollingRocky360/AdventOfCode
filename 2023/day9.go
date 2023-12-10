package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func reverse(levels [][]int) [][]int {
	rev := make([][]int, len(levels))
	copy(rev, levels)

	for i := 0; i < len(rev)/2; i++ {
		rev[i], rev[len(rev)-i-1] = rev[len(rev)-i-1], rev[i]
	}

	return rev
}

func extrapolate(nums []int) int {
	levels := [][]int{nums}
	currLevel := levels[0]

	for !allZeros(currLevel) {
		newLevel := make([]int, len(currLevel)-1)

		for i := 0; i < len(currLevel)-1; i++ {
			newLevel[i] = currLevel[i+1] - currLevel[i]
		}

		levels = append(levels, newLevel)
		currLevel = newLevel
	}

	intermediate := 0
	for _, level := range reverse(levels) {
		// intermediate += level[len(level)-1]  // part 1
		intermediate = level[0] - intermediate
	}

	return intermediate
}

func day9() {
	inputFile, err := os.Open("inputs/day9.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	total := 0
	for scanner.Scan() {
		nums := []int{}
		for _, numStr := range strings.Fields(scanner.Text()) {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		total += extrapolate(nums)
	}

	fmt.Println(total)
}
