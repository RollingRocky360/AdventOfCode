package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var MAX_BLINKS = 25

var cache = map[string]int{}

func getCount(numText string, iter int) int {
	if iter == MAX_BLINKS+1 {
		return 1
	}

	state := numText + fmt.Sprintf(",%d", iter)

	if cached := cache[state]; cached != 0 {
		return cached
	}

	if numText == "0" {
		c := getCount("1", iter+1)
		cache[state] = c
		return c
	}

	if len(numText)%2 == 0 {
		l := len(numText)
		leftText := numText[0 : l/2]
		rightNum, _ := strconv.Atoi(numText[l/2:])
		rightText := strconv.Itoa(rightNum)
		c := getCount(leftText, iter+1) + getCount(rightText, iter+1)
		cache[state] = c
		return c
	}

	num, _ := strconv.Atoi(numText)
	numTextMul := strconv.Itoa(num * 2024)
	c := getCount(numTextMul, iter+1)
	cache[state] = c
	return c
}

func Part1() {
	stones := parseInput()
	MAX_BLINKS = 25

	count := 0
	for _, stone := range stones {
		count += getCount(stone, 1)
	}

	log.Println("part 1:", count)
}

func Part2() {
	stones := parseInput()
	MAX_BLINKS = 75

	for k := range cache {
		delete(cache, k)
	}

	count := 0
	for _, stone := range stones {
		count += getCount(stone, 1)
	}

	log.Println("part 2:", count)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() []string {
	input, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	text := scanner.Text()

	return strings.Fields(text)
}
