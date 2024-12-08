package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result   int64
	operands []int64
}

func isCorrect(result int64, intermediate int64, operands []int64) bool {
	if len(operands) == 0 {
		return result == intermediate
	}
	return isCorrect(result, intermediate*operands[0], operands[1:]) ||
		isCorrect(result, intermediate+operands[0], operands[1:])
}

func Part1() {
	equations := parseInput()

	sumCorrect := int64(0)
	for _, eq := range equations {
		if isCorrect(eq.result, eq.operands[0], eq.operands[1:]) {
			sumCorrect += eq.result
		}
	}

	log.Println("part 1:", sumCorrect)
}

func isCorrectWithConcat(result int64, intermediate int64, operands []int64) bool {
	if len(operands) == 0 {
		return result == intermediate
	}

	if isCorrectWithConcat(result, intermediate*operands[0], operands[1:]) ||
		isCorrectWithConcat(result, intermediate+operands[0], operands[1:]) {
		return true
	}

	concatedText := fmt.Sprintf("%d", intermediate) + fmt.Sprintf("%d", operands[0])
	conated, _ := strconv.ParseInt(concatedText, 10, 64)
	return isCorrectWithConcat(result, conated, operands[1:])
}

func Part2() {
	equations := parseInput()

	sumCorrect := int64(0)
	for _, eq := range equations {
		if isCorrectWithConcat(eq.result, eq.operands[0], eq.operands[1:]) {
			sumCorrect += eq.result
		}
	}

	log.Println("part 2:", sumCorrect)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() []equation {
	input, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	equations := []equation{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		var result int64
		var operands []int64

		splits := strings.Split(line, ": ")
		resultText, operandsText := splits[0], splits[1]
		result, _ = strconv.ParseInt(resultText, 10, 64)

		if resultText != fmt.Sprintf("%d", result) {
			fmt.Println(resultText)
		}

		for _, operandText := range strings.Fields(operandsText) {
			o, _ := strconv.ParseInt(operandText, 10, 64)
			operands = append(operands, o)
		}

		equations = append(equations, equation{result, operands})
	}
	return equations
}
