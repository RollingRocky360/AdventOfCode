package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var addrRegex *regexp.Regexp = regexp.MustCompile(`\[|\]`)

func checkTLSValidity(addr string) int {
	res := addrRegex.Split(addr, -1)

	valid := false
	aba := map[string]bool{}
	for i := 0; i < len(res); i += 2 {
		part := res[i]
		for ind := 0; ind+2 < len(part); ind++ {
			if part[ind] == part[ind+2] && part[ind] != part[ind+1] {
				valid = true
				aba[part[ind:ind+3]] = true
			}
		}
	}

	if !valid {
		return 0
	}

	for i := 1; i < len(res); i += 2 {
		part := res[i]
		for ind := 0; ind+2 < len(part); ind++ {
			if part[ind] == part[ind+2] && part[ind] != part[ind+1] {
				bab := string(part[ind+1]) + string(part[ind]) + string(part[ind+1])
				if aba[bab] {
					return 1
				}
			}
		}
	}

	return 0
}

func day7() {
	inputFile, _ := os.Open("inputs/Day7.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	cnt := 0
	for scanner.Scan() {
		cnt += checkTLSValidity(scanner.Text())
	}

	fmt.Println(cnt)
}
