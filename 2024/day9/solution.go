package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FREE_BLOCK = -1

func filledSlice(length int, element int) []int {
	slice := make([]int, length)

	if element == 0 {
		return slice
	}

	for i := range slice {
		slice[i] = element
	}
	return slice
}

func getNextFreeBlock(unzipped []int, i int) int {
	for i < len(unzipped) && unzipped[i] != FREE_BLOCK {
		i++
	}
	return i
}

func getNextFileBlock(unzipped []int, i int) int {
	for i >= 0 && unzipped[i] == FREE_BLOCK {
		i--
	}
	return i
}

func getFileStartBlock(unzipped []int, i int) int {
	fileId := unzipped[i]
	for i >= 0 && unzipped[i] == fileId {
		i--
	}
	return i + 1
}

func getFreeEndBlock(unzipped []int, i int) int {
	for i < len(unzipped) && unzipped[i] == FREE_BLOCK {
		i++
	}
	return i - 1
}

func getLeftMostFreeSpanThatFits(unzipped []int, boundary, length int) (int, int, bool) {
	i, j := -1, -1
	for {
		i = getNextFreeBlock(unzipped, i+1)
		j = getFreeEndBlock(unzipped, i)
		if i > boundary {
			break
		}
		if j-i+1 >= length {
			return i, j, true
		}
	}
	return -1, -1, false
}

func computeChecksum(unzipped []int) int {
	var checksum int

	for pos, block := range unzipped {
		if block == FREE_BLOCK {
			continue
		}
		checksum += pos * block
	}

	return checksum
}

func Display(unzipped []int) {
	for _, block := range unzipped {
		if block == -1 {
			fmt.Print(". ")
		} else {
			fmt.Printf("%d ", block)
		}
	}
	fmt.Println()
}

func Part1() {
	mapping := parseInput()
	unzipped := []int{}

	for i, blockSize := range mapping {
		filler := -1 // free block
		if i%2 == 0 {
			filler = i / 2 // file ID
		}

		intermediate := filledSlice(blockSize, filler)
		unzipped = append(unzipped, intermediate...)
	}

	freePointer := 0
	filePointer := len(unzipped) - 1

	for {
		freePointer = getNextFreeBlock(unzipped, freePointer)
		filePointer = getNextFileBlock(unzipped, filePointer)
		if freePointer >= filePointer {
			break
		}
		unzipped[freePointer], unzipped[filePointer] = unzipped[filePointer], unzipped[freePointer]
	}

	checksum := computeChecksum(unzipped)
	log.Println("part 1:", checksum)
}

func Part2() {
	mapping := parseInput()
	unzipped := []int{}

	for i, blockSize := range mapping {
		filler := FREE_BLOCK
		if i%2 == 0 {
			filler = i / 2 // file ID
		}

		intermediate := filledSlice(blockSize, filler)
		unzipped = append(unzipped, intermediate...)
	}

	filePointer := -1
	fileStartPointer := len(unzipped)
	// Display(unzipped)

	for {
		filePointer = getNextFileBlock(unzipped, fileStartPointer-1)
		if filePointer < 0 {
			break
		}
		fileStartPointer = getFileStartBlock(unzipped, filePointer)

		length := filePointer - fileStartPointer + 1

		freeStart, freeEnd, fits := getLeftMostFreeSpanThatFits(unzipped, fileStartPointer-1, length)
		if !fits {
			continue
		}

		copy(unzipped[freeStart:freeEnd+1], unzipped[fileStartPointer:filePointer+1])
		for i := fileStartPointer; i <= filePointer; i++ {
			unzipped[i] = -1
		}
		// Display(unzipped)
	}

	checksum := computeChecksum(unzipped)
	log.Println("part 2:", checksum)
}

func Run() {
	Part1()
	Part2()
}

func parseInput() []int {
	input, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	mappingBytes := scanner.Bytes()

	mapping := []int{}
	for _, b := range mappingBytes {
		mapping = append(mapping, int(b)-48)
	}

	return mapping
}
