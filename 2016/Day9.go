package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var markerRegex *regexp.Regexp = regexp.MustCompile(`\((\d+)x(\d+)\)`)

func computeLength(text string) int {
	if len(text) == 0 {
		return 0
	}

	count := 0
	for len(text) > 0 {
		pos := markerRegex.FindStringSubmatchIndex(text)
		if len(pos) == 0 {
			break
		}

		markerStart, markerEnd := pos[0], pos[1]
		ogLength, _ := strconv.Atoi(text[pos[2]:pos[3]])
		reps, _ := strconv.Atoi(text[pos[4]:pos[5]])

		innerDecompLength := computeLength(text[markerEnd : markerEnd+ogLength])

		count += markerStart + innerDecompLength*reps
		text = text[markerEnd+ogLength:]
	}

	count += len(text)
	return count
}

func day9() {
	inputFile, _ := os.Open("inputs/Day9.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println(computeLength(text))
}
