package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var gameRegex = regexp.MustCompile(`(\d+) (red|green|blue)`)
var constraints = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isValidGame(line string) bool {
	for _, match := range gameRegex.FindAllStringSubmatch(line, -1) {
		foundBalls, _ := strconv.Atoi(match[1])
		if constraints[match[2]] < foundBalls {
			return false
		}
	}
	return true
}

func power(line string) int {
	maxi := map[string]int{
		"red": 0, "green": 0, "blue": 0,
	}

	for _, match := range gameRegex.FindAllStringSubmatch(line, -1) {
		foundBalls, _ := strconv.Atoi(match[1])
		maxi[match[2]] = maxInt(maxi[match[2]], foundBalls)
	}

	return maxi["red"] * maxi["green"] * maxi["blue"]
}

func day2() {
	inputFile, err := os.Open("inputs/day2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	gameNumber := 1
	ans := 0
	for scanner.Scan() {
		ans += power(scanner.Text())
		gameNumber += 1
	}

	fmt.Println(ans)
}
