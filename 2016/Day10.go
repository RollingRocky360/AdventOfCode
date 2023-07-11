package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var botConfig [300][2]int = [300][2]int{}
var botInstr [300][4]int = [300][4]int{}
var outputs [300][]int = [300][]int{}

func assignToBot(bot, val int) {
	low := botConfig[bot][0]
	if low == 0 {
		botConfig[bot][0] = val
		return
	}

	if val > low {
		botConfig[bot][1] = val
	} else {
		botConfig[bot][1] = botConfig[bot][0]
		botConfig[bot][0] = val
	}
}

func distribute(bot int) {
	instr, config := botInstr[bot][:], botConfig[bot][:]
	if config[0] == 0 || config[1] == 0 {
		return
	}

	if config[0] == 17 && config[1] == 61 {
		fmt.Println("bot comparing 61 with 17:", bot)
	}

	if instr[2] == 1 {
		outputs[instr[0]] = append(outputs[instr[0]], config[0])
	} else {
		assignToBot(instr[0], config[0])
	}
	config[0] = 0

	if instr[3] == 1 {
		outputs[instr[1]] = append(outputs[instr[1]], config[1])
	} else {
		assignToBot(instr[1], config[1])
	}
	config[1] = 0

	if instr[2] != 1 {
		distribute(instr[0])
	}
	if instr[3] != 1 {
		distribute(instr[1])
	}
}

func day10() {
	inputFile, _ := os.Open("inputs/Day10.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	botInitRegex := regexp.MustCompile(`^value (\d+) goes to bot (\d+)$`)
	botInstrRegex := regexp.MustCompile(
		`^bot (\d+) gives low to (output|bot) (\d+) and high to (output|bot) (\d+)$`,
	)

	maxBotNumber := 0

	for scanner.Scan() {
		instr := scanner.Text()

		if sub := botInstrRegex.FindStringSubmatch(instr); len(sub) > 0 {
			var (
				lowToOP  int
				highToOP int
			)

			bot, _ := strconv.Atoi(sub[1])
			if bot > maxBotNumber {
				maxBotNumber = bot
			}

			lowTo, _ := strconv.Atoi(sub[3])
			if sub[2] == "output" {
				lowToOP = 1
			} else if lowTo > maxBotNumber {
				maxBotNumber = lowTo
			}

			highTo, _ := strconv.Atoi(sub[5])
			if sub[4] == "output" {
				highToOP = 1
			} else if highTo > maxBotNumber {
				maxBotNumber = highTo
			}

			botInstr[bot] = [4]int{lowTo, highTo, lowToOP, highToOP}

		} else if sub := botInitRegex.FindStringSubmatch(instr); len(sub) > 0 {
			val, _ := strconv.Atoi(sub[1])
			toBot, _ := strconv.Atoi(sub[2])
			if toBot > maxBotNumber {
				maxBotNumber = toBot
			}
			assignToBot(toBot, val)
		}
	}

	for bot := 0; bot <= maxBotNumber; bot++ {
		distribute(bot)
	}

	product := 1
	for _, binContent := range outputs[:3] {
		for _, val := range binContent {
			product *= val
		}
	}
	fmt.Println("Multiplaying first three output bins:", product)
}
