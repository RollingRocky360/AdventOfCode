package day3

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func Part1() {
	mem := getInput()
	uncorruptedMuls := regexp.
		MustCompile(`mul\(\d{1,3},\d{1,3}\)`).
		FindAllString(mem, -1)

	total := 0
	for _, mul := range uncorruptedMuls {
		var a, b int
		fmt.Sscanf(mul, "mul(%d,%d)", &a, &b)
		total += a * b
	}

	log.Println("part 1:", total)
}

func Part2() {
	mem := getInput()
	instrs := regexp.
		MustCompile(`mul\(\d{1,3},\d{1,3}\)|do(n't)?\(\)`).
		FindAllString(mem, -1)

	total := 0
	enabled := true

	for _, instr := range instrs {
		switch instr {
		case "don't()":
			enabled = false
		case "do()":
			enabled = true
		default:
			if !enabled {
				continue
			}
			var a, b int
			fmt.Sscanf(instr, "mul(%d,%d)", &a, &b)
			total += a * b
		}
	}

	log.Println("part 2:", total)
}

func getInput() string {
	input, err := os.Open("day3/test-input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	text, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(text)
}

func Run() {
	Part1()
	Part2()
}
