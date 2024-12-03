package day2

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(a int) int {
	return max(a, -a)
}

func parseInput() [][]int {
	input, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer input.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		report, err := parseReportText(scanner.Text())
		if err != nil {
			log.Fatalln("corrupted input")
		}

		reports = append(reports, report)
	}

	return reports
}

func parseReportText(reportText string) ([]int, error) {
	levels := []int{}
	for _, levelText := range strings.Fields(reportText) {
		level, err := strconv.Atoi(levelText)
		if err != nil {
			return nil, err
		}
		levels = append(levels, level)
	}
	return levels, nil
}

func isReportSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	prevDiff := report[1] - report[0]
	if abs(prevDiff) < 1 || abs(prevDiff) > 3 {
		return false
	}

	for i := 2; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if abs(diff) < 1 || abs(diff) > 3 || prevDiff^diff < 0 {
			return false
		}

		prevDiff = diff
	}

	return true
}

func isReportSafeWithDamp(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	if isReportSafe(report[1:]) {
		return true
	}

	for i := 2; i < len(report); i++ {
		droppedOne := slices.Concat(report[0:i-1], report[i:])
		if isReportSafe(droppedOne) {
			return true
		}
	}

	if isReportSafe(report[:len(report)-1]) {
		return true
	}

	return false
}

func Part1() {
	reports := parseInput()

	safeCount := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeCount++
		}
	}

	log.Println("part 1:", safeCount)
}

func Part2() {
	reports := parseInput()

	safeCount := 0
	for _, report := range reports {
		if !isReportSafe(report) && !isReportSafeWithDamp(report) {
			continue
		}
		safeCount++
	}

	log.Println("part 2:", safeCount)
}

func Run() {
	Part1()
	Part2()
}
