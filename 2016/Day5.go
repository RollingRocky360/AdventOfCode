package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func getEncodedMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func door1(text string) (string, bool) {
	encoded := getEncodedMd5(text)
	return string(encoded[5]), encoded[:5] == "00000"
}

func door2(text string) (int, string, bool) {
	encoded := getEncodedMd5(text)
	pos, err := strconv.Atoi(string(encoded[5]))
	return pos, string(encoded[6]), encoded[:5] == "00000" && pos < 8 && err == nil
}

func day5() {
	text := "abbhdwsy"
	cipher := [8]string{}

	for i, cnt := 0, 0; i < 100000000; i++ {
		pos, char, ok := door2(text + strconv.Itoa(i))
		if !ok || cipher[pos] != "" {
			continue
		}

		cipher[pos] = char
		cnt++
	}

	answer := ""
	for _, ch := range cipher {
		answer += ch
	}

	fmt.Println(cipher, answer)
}
