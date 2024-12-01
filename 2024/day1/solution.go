package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	list1, list2 := parseInput()

	slices.Sort(list1)
	slices.Sort(list2)

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		totalDistance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	log.Println("part 1:", totalDistance)
}

func Part2() {
	list1, list2 := parseInput()
	countMap := map[int]int{}

	for _, num := range list2 {
		countMap[num]++
	}

	similarity := 0
	for _, num := range list1 {
		similarity += num * countMap[num]
	}

	log.Println("part 2:", similarity)
}

func Run() {
	Part1()
	Part2()
}

// parses input file and returns the two lists of location IDs
func parseInput() ([]int, []int) {
	input, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer input.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		loc1, _ := strconv.Atoi(line[0])
		list1 = append(list1, loc1)
		loc2, _ := strconv.Atoi(line[1])
		list2 = append(list2, loc2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err.Error())
	}

	if len(list1) != len(list2) {
		log.Fatalln("input is corrupted: length of the lists are not the same")
	}

	return list1, list2
}
