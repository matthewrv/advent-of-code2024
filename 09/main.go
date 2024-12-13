package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	space := readInput("input.txt")

	// fmt.Print("Initial space:      ")
	// printSpace(space)

	result := defragmentPart2(space)

	// fmt.Print("Defragmented space: ")
	// printSpace(space)

	fmt.Println("Result:", result)
}

// common

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

// part 1

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

// part 2

func defragmentPart2(space []*int) (checksum int) {
	r := len(space) - 1
	for r > 0 {
		// find file block
		for space[r] == nil {
			r--
		}

		size := 0
		fileId := space[r]
		for r-size > 0 && space[r-size] == fileId {
			size++
		}

		// find free space
		l, available := 0, 0
		for l < r && available != size {
			// find start of free space
			for space[l] != nil {
				l++
			}

			// check if size is enough
			for available != size && l+available < len(space) && space[l+available] == nil {
				available++
			}

			if available != size {
				l = l + available
				available = 0
			}
		}

		if l < r && available == size {
			// swap
			for i := 0; i < size; i++ {
				space[l+i] = fileId
				space[r-i] = nil
			}
		}
		r = r - size
	}

	// calc checksum
	for idx, blockId := range space {
		if blockId != nil {
			checksum += idx * *blockId
		}
	}

	return checksum
}
