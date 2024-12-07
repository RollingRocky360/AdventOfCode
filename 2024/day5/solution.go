package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isInOrder(ruleMap map[int]map[int]bool, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		left, right := update[i], update[i+1]
		if !ruleMap[left][right] {
			return false
		}
	}
	return true
}

func intersectionCount(a, b map[int]bool) int {
	count := 0
	for k := range a {
		if b[k] {
			count++
		}
	}
	return count
}

func findMiddleAfterCorrection(ruleMap map[int]map[int]bool, update []int) int {
	length := len(update)

	updateSet := map[int]bool{}
	for _, page := range update {
		updateSet[page] = true
	}

	for _, page := range update {
		if intersectionCount(ruleMap[page], updateSet) == length/2 {
			return page
		}
	}

	return 0
}

func Part1() {
	ruleMap, updates := parseInput()

	sumMiddle := 0
	for _, update := range updates {
		length := len(update)
		if isInOrder(ruleMap, update) {
			sumMiddle += update[length/2]
		}
	}

	log.Println("part 1:", sumMiddle)
}

func Part2() {
	ruleMap, updates := parseInput()

	sumMiddle := 0
	for _, update := range updates {
		if !isInOrder(ruleMap, update) {
			sumMiddle += findMiddleAfterCorrection(ruleMap, update)
		}
	}

	log.Println("part 2:", sumMiddle)
}

func Run() {
	Part1()
	Part2()
}

type rule struct {
	left  int
	right int
}

func parseInput() (map[int]map[int]bool, [][]int) {
	input, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var rules []rule
	for scanner.Scan() && scanner.Text() != "" {
		r := rule{}
		fmt.Sscanf(scanner.Text(), "%d|%d", &r.left, &r.right)
		rules = append(rules, r)
	}

	var updates [][]int
	for scanner.Scan() {
		update := []int{}
		for _, page := range strings.Split(scanner.Text(), ",") {
			pageNumber, _ := strconv.Atoi(page)
			update = append(update, pageNumber)
		}
		updates = append(updates, update)
	}

	ruleMap := map[int]map[int]bool{}

	for _, r := range rules {
		if ruleMap[r.left] == nil {
			ruleMap[r.left] = map[int]bool{}
		}
		ruleMap[r.left][r.right] = true
	}

	return ruleMap, updates
}
