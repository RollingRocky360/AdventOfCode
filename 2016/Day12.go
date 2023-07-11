package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instr struct {
	code      int
	arg1      int
	isArg1Reg bool
	arg2      int
	isArg2Reg bool
}

const (
	cpy int = iota
	inc
	dec
	jnz
)

var registers [4]int = [4]int{0, 0, 1, 0}

var instructions []instr = []instr{}

func parseInstructions() {
	inputFile, _ := os.Open("inputs/Day12.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		sub := strings.Split(scanner.Text(), " ")
		newInstr := instr{}

		switch sub[0] {
		case "cpy":
			newInstr.code = cpy
		case "inc":
			newInstr.code = inc
		case "dec":
			newInstr.code = dec
		case "jnz":
			newInstr.code = jnz
		}

		if reg := sub[1][0]; reg >= 97 && reg <= 102 {
			newInstr.arg1 = int(reg) - 97
			newInstr.isArg1Reg = true
		} else {
			newInstr.arg1, _ = strconv.Atoi(sub[1])
		}

		if len(sub) > 2 {
			if reg := sub[2][0]; reg >= 97 && reg <= 102 {
				newInstr.arg2 = int(reg) - 97
				newInstr.isArg2Reg = true
			} else {
				newInstr.arg2, _ = strconv.Atoi(sub[2])
			}
		}

		instructions = append(instructions, newInstr)
	}
}

func day12() {
	parseInstructions()
	mem := 0

	for mem < len(instructions) {
		i := instructions[mem]
		switch i.code {
		case cpy:
			val := i.arg1
			if i.isArg1Reg {
				val = registers[val]
			}
			dest := i.arg2
			registers[dest] = val
		case inc:
			registers[i.arg1]++
		case dec:
			registers[i.arg1]--
		case jnz:
			val := i.arg1
			if i.isArg1Reg {
				val = registers[val]
			}
			if val != 0 {
				mem += i.arg2
				continue
			}
		}

		mem++
	}

	fmt.Println(registers)
}
