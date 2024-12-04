package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// since puzzle contains only ascii chars - byte array is enough
	searchWord := []byte("XMAS")

	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := bytes.Split(fileContent, []byte("\n"))
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

	var wordCount int
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
	fmt.Printf("Total word count: %d\n", wordCount)
}

func checkWord(content [][]byte, searchWord []byte, startLine int, startCol int, vector [2]int) int {
	totalLines := len(content)
	totalColumns := len(content[0])

	for idx, char := range searchWord {
		line, col := startLine+idx*vector[0], startCol+idx*vector[1]
		if line < 0 || col < 0 || line >= totalLines || col >= totalColumns {
			return 0
		}
		if content[line][col] != char {
			return 0
		}
	}
	return 1
}
