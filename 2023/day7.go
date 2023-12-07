package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var signatureOrdering = map[string]int{
	"5":     0,
	"14":    1,
	"23":    2,
	"113":   3,
	"122":   4,
	"1112":  5,
	"11111": 6,
}

var cardOrdering = map[byte]int{
	'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7,
	'T': 8, 'J': -1, 'Q': 10, 'K': 11, 'A': 12,
	// Chane J:-1 to to J:9 for part 1
}

type HandAndBid struct {
	hand string
	bid  int
}

func handType(hand string) int {
	freq := map[rune]int{}
	for _, r := range hand {
		freq[r]++
	}

	// Omit this if-block for part 1
	if freq['J'] != 0 {
		maxCard, maxCount := 'J', 0
		for k, v := range freq {
			if v > maxCount && k != 'J' {
				maxCard, maxCount = k, v
			}
		}
		delete(freq, maxCard)
		freq['J'] += maxCount
	}

	counts := []int{}
	for _, v := range freq {
		counts = append(counts, v)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] < counts[j]
	})

	signature := ""
	for _, count := range counts {
		signature += strconv.Itoa(count)
	}

	return signatureOrdering[signature]
}

func handLessThan(hand1, hand2 string) bool {
	for i := 0; i < 5; i++ {
		o1, o2 := cardOrdering[hand1[i]], cardOrdering[hand2[i]]
		if o1 != o2 {
			return o1 < o2
		}
	}
	return false
}

func day7() {
	inputFile, err := os.Open("inputs/day7.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inputFile.Close()

	hAndB := []HandAndBid{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		hand := f[0]
		bid, _ := strconv.Atoi(f[1])
		hAndB = append(hAndB, HandAndBid{hand, bid})
	}

	sort.Slice(hAndB, func(i, j int) bool {
		iType := handType(hAndB[i].hand)
		jType := handType(hAndB[j].hand)

		if iType == jType {
			return handLessThan(hAndB[i].hand, hAndB[j].hand)
		}

		return iType > jType
	})

	total := 0
	for ind, hand := range hAndB {
		total += (ind + 1) * hand.bid
	}

	fmt.Println(total)
}
