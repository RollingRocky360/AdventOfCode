package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var seedsRegex = regexp.MustCompile(`^seeds: ((?:\d+ ?)+)$`)
var mappingRegex = regexp.MustCompile(`^(\d+) (\d+) (\d+)$`)

type mapping struct {
	destination int
	source      int
	quantity    int
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func seedToLocation(seed int, titles []string, mappings map[string][]mapping) int {
	intermediate := seed
	for _, title := range titles {
		for _, m := range mappings[title] {
			if intermediate >= m.source && intermediate < m.source+m.quantity {
				intermediate = m.destination + intermediate - m.source
				break
			}
		}
	}
	return intermediate
}

func day5() {
	inputFile, err := os.Open("inputs/day5.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// Scan Seeds
	scanner.Scan()
	seedString := seedsRegex.FindStringSubmatch(scanner.Text())[1]
	seeds := []int{}
	for _, num := range strings.Split(seedString, " ") {
		intNum, _ := strconv.Atoi(num)
		seeds = append(seeds, intNum)
	}

	mappings := map[string][]mapping{}
	titles := []string{}
	scanner.Scan()

	for scanner.Scan() {
		title := scanner.Text()
		titles = append(titles, title)

		newMapping := []mapping{}

		for scanner.Scan() && scanner.Text() != "" {
			m := mappingRegex.FindStringSubmatch(scanner.Text())
			d, _ := strconv.Atoi(m[1])
			s, _ := strconv.Atoi(m[2])
			q, _ := strconv.Atoi(m[3])
			newMapping = append(newMapping, mapping{destination: d, source: s, quantity: q})
		}

		mappings[title] = newMapping
		sort.Slice(mappings[title], func(i, j int) bool {
			return mappings[title][i].source < mappings[title][j].source
		})
	}

	// part 1
	minimal := math.MaxInt
	for _, seed := range seeds {
		minimal = minInt(minimal, seedToLocation(seed, titles, mappings))
	}
	fmt.Println(minimal)

	// part 2 - bruteforce
	minimal = math.MaxInt
	start := time.Now()
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			minimal = minInt(minimal, seedToLocation(seed, titles, mappings))
		}
	}

	end := time.Since(start)
	fmt.Println(minimal, "- took", end, "seconds")

	// part 2 - suboptimal
	// todo
}
