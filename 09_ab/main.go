package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	space := readInput("input.txt")

	// fmt.Print("Initial space: ")
	// printSpace(space)

	result := defragment(space)

	// fmt.Print("Defragmented space: ")
	// printSpace(space)

	fmt.Println("Result:", result)
}

// part 1

func readInput(fileName string) (space []*int) {
	content, _ := os.ReadFile(fileName)

	var idx, isFree = 0, false
	for _, char := range content {
		size, _ := strconv.Atoi(string(char))
		var val *int
		if !isFree {
			val = new(int)
			*val = idx
			idx++
		}
		for i := 0; i < size; i++ {
			space = append(space, val)
		}

		isFree = !isFree
	}

	return space
}

func defragment(space []*int) (checksum int) {
	l, r := 0, len(space)-1
	for l < r {
		for space[r] == nil {
			r--
		}
		for space[l] != nil {
			l++
		}
		if l < r {
			space[l], space[r] = space[r], space[l]
		}
	}

	for idx, blockId := range space {
		if blockId != nil {
			checksum += idx * *blockId
		}
	}

	return checksum
}

func printSpace(space []*int) {
	for _, ref := range space {
		if ref != nil {
			fmt.Print(*ref, " ")
		} else {
			fmt.Print(".", " ")
		}
	}
	fmt.Println()
}
