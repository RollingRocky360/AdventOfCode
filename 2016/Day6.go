package main

import (
	"bufio"
	"fmt"
	"os"
)

func day6() {
	inputFile, _ := os.Open("inputs/Day6.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	messages := []string{}
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	freqs := make(map[string]int)
	mostFreq, cnt := "a", 1000
	res := ""

	for i := 0; i < len(messages[0]); i++ {

		for _, msg := range messages {
			freqs[string(msg[i])]++
		}

		for char, freq := range freqs {
			if freq < cnt {
				mostFreq, cnt = char, freq
			}
		}

		res += mostFreq
		cnt = 1000
		for k := range freqs {
			delete(freqs, k)
		}

	}

	fmt.Println(res)
}
