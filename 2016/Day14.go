package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

var salt = "ahsbgdzn"
var buffer = []string{}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	encoded := hex.EncodeToString(hash[:])
	for i := 0; i < 2016; i++ {
		hash = md5.Sum([]byte(encoded))
		encoded = hex.EncodeToString(hash[:])
	}
	return encoded
}

func hasRep(text string, reps int) (byte, bool) {
	for i := 0; i+reps <= len(text); i++ {
		if strings.Count(text[i:i+reps], string(text[i])) == reps {
			return text[i], true
		}
	}
	return 0, false
}

func hasRepSpecific(text string, reps int, sub byte) bool {
	for i := 0; i+reps <= len(text); i++ {
		if text[i] == sub && strings.Count(text[i:i+reps], string(text[i])) == reps {
			return true
		}
	}
	return false
}

func day14() {
	count, ind := 0, 0

	for ind = 0; ind < 1000; ind++ {
		buffer = append(buffer, MD5(salt+strconv.Itoa(ind)))
	}

	for ind = 0; ind < 50000 && count < 64; ind++ {
		hash := buffer[0]
		buffer = buffer[1:]
		buffer = append(buffer, MD5(salt+strconv.Itoa(ind+1000)))

		r, ok := hasRep(hash, 3)
		if !ok {
			continue
		}

		for j := 0; j < 1000; j++ {
			if hasRepSpecific(buffer[j], 5, r) {
				count++
			}
		}
	}

	fmt.Println(ind - 1)
}
