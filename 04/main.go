package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := bytes.Split(fileContent, []byte("\n"))

	// totalCount := xmasPuzzle(content)
	totalCount := xMasPuzzle(content)

	fmt.Printf("Total word count: %d\n", totalCount)
}

// part 1

func xmasPuzzle(content [][]byte) (wordCount int) {
	// since puzzle contains only ascii chars - byte array is enough
	searchWord := []byte("XMAS")

	vectors := [8][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	for lineIdx, line := range content {
		for colIdx, char := range line {
			// not start of the word
			if char != searchWord[0] {
				continue
			}

			for _, vector := range vectors {
				wordCount += checkWord(content, searchWord, lineIdx, colIdx, vector)
			}

		}
	}

	return wordCount
}

func checkWord(content [][]byte, searchWord []byte, startLine int, startCol int, vector [2]int) int {
	totalLines := len(content)
	totalColumns := len(content[0])

	for idx, char := range searchWord {
		line, col := startLine+idx*vector[0], startCol+idx*vector[1]
		// check slice boundaries
		if line < 0 || col < 0 || line >= totalLines || col >= totalColumns {
			return 0
		}
		if content[line][col] != char {
			return 0
		}
	}
	return 1
}

// part 2

func xMasPuzzle(content [][]byte) (xMasCount int) {
	// searchWord := []byte("MAS")

	for lineIdx, line := range content {
		for colIdx, char := range line {
			if char == 'A' {
				xMasCount += checkXMas(content, lineIdx, colIdx)
			}
		}
	}

	return xMasCount
}

func checkXMas(content [][]byte, lineIdx int, colIdx int) int {
	// if A is in first or last column or row - it can not be X-Mas
	if lineIdx == 0 || colIdx == 0 || lineIdx == len(content)-1 || colIdx == len(content[0])-1 {
		return 0
	}

	validWords := [2][]byte{
		[]byte("MAS"),
		[]byte("SAM"),
	}

	// check first diagonal
	word1 := []byte{content[lineIdx-1][colIdx-1], 'A', content[lineIdx+1][colIdx+1]}
	if !bytes.Equal(word1, validWords[0]) && !bytes.Equal(word1, validWords[1]) {
		return 0
	}

	// check second diagonal
	word2 := []byte{content[lineIdx-1][colIdx+1], 'A', content[lineIdx+1][colIdx-1]}
	if !bytes.Equal(word2, validWords[0]) && !bytes.Equal(word2, validWords[1]) {
		return 0
	}

	return 1
}
