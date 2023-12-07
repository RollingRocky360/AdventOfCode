package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func quadraticInequalitySolution(b, c float64) (float64, float64) {
	lowerBound := math.Floor((b - math.Sqrt(b*b-4*c)) / 2)
	upperBound := math.Ceil((b + math.Sqrt(b*b-4*c)) / 2)
	return lowerBound, upperBound
}

func day6() {
	inputFile, err := os.Open("inputs/day6.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	bStr, cStr := "", ""

	scanner.Scan()
	times := []float64{}
	for _, num := range strings.Fields(scanner.Text())[1:] {
		x, _ := strconv.Atoi(num)
		bStr += num
		times = append(times, float64(x))
	}

	scanner.Scan()
	distances := []float64{}
	for _, num := range strings.Fields(scanner.Text())[1:] {
		x, _ := strconv.Atoi(num)
		cStr += num
		distances = append(distances, float64(x))
	}

	// part 1
	ans := 1
	for i := 0; i < len(times); i++ {
		b, c := times[i], distances[i]
		lowerBound, upperBound := quadraticInequalitySolution(b, c)
		ans *= int(upperBound - lowerBound - 1)
	}

	fmt.Println("Part 1 -", ans)

	// part 2
	b, _ := strconv.Atoi(bStr)
	c, _ := strconv.Atoi(cStr)
	fmt.Println(b, c)
	lowerBound, upperBound := quadraticInequalitySolution(float64(b), float64(c))
	fmt.Println("Part 2 -", upperBound-lowerBound-1)
}
