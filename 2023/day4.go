package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func contains[T comparable](container []T, element T) bool {
	for _, i := range container {
		if i == element {
			return true
		}
	}
	return false
}

func day4() {
	inputFile, err := os.Open("inputs/day4.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	splitRegex := regexp.MustCompile(`^Card +(\d+):((?: +\d+)*) \|((?: +\d+)*)$`)
	numberRegex := regexp.MustCompile(`\d+`)

	counts := map[int]int{}
	totalCards := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		split := splitRegex.FindStringSubmatch(scanner.Text())

		winning := numberRegex.FindAllString(split[2], -1)
		having := numberRegex.FindAllString(split[3], -1)
		cardNo, _ := strconv.Atoi(split[1])

		matchCount := 0
		for _, h := range having {
			if contains(winning, h) {
				matchCount++
			}
		}

		counts[cardNo]++
		for c := cardNo + 1; c <= cardNo+matchCount; c++ {
			counts[c] += counts[cardNo]
		}

		totalCards += counts[cardNo]
	}

	fmt.Println(totalCards)
}
