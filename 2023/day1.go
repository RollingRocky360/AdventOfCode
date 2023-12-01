package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var value = `one|two|three|four|five|six|seven|eight|nine|\d`
var mainRe = regexp.MustCompile(fmt.Sprintf(`^[a-z]*?(%s).*(%s)[a-z]*$`, value, value))
var backupRe = regexp.MustCompile(value)

var numMap = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
	"six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func textToNum(num string) string {
	if len(num) == 1 {
		return num
	}
	return numMap[num]
}

func extractNumber(line string) int {
	res := mainRe.FindStringSubmatch(line)
	var d1, d2 string
	if len(res) == 0 {
		d1 = textToNum(backupRe.FindString(line))
		d2 = d1
	} else {
		d1 = textToNum(res[1])
		d2 = textToNum(res[2])
	}

	num, _ := strconv.Atoi(d1 + d2)
	return num

}

func day1() {
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += extractNumber(scanner.Text())
	}

	fmt.Println(total)
}
