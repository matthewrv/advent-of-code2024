package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	size, obstructionsMap, position, direction := readInput("input.txt")
	visited := countVisitedPositions(size, obstructionsMap, position, direction)
	fmt.Printf("Result: %d\n", visited)
}

// helper struct

type vector struct {
	x int
	y int
}

func (v1 *vector) Sum(v2 *vector) vector {
	return vector{v1.x + v2.x, v1.y + v2.y}
}

// read input

func readInput(fileName string) (size vector, obstructionsMap map[vector]bool, position vector, direction vector) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	line := 0
	obstructionsMap = map[vector]bool{}
	for scanner.Scan() {
		row := scanner.Text()
		for col, token := range row {
			switch token {
			case '#':
				obstructionsMap[vector{col, line}] = true
			case '^':
				position = vector{col, line}
				// axis are assumed to increase in direction of increasing index of array
				direction = vector{0, -1}
			}
		}

		size.x = len(row)
		line++
	}

	size.y = line + 1

	return size, obstructionsMap, position, direction
}

// part 1

func countVisitedPositions(size vector, obstructionsMap map[vector]bool, position vector, direction vector) int {
	visitedPositions := map[vector]bool{}
	visitedPositions[position] = true

	gridSizeX := size.x
	gridSizeY := size.y

	rotationRules := map[vector]vector{
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{0, 1}:  {-1, 0},
		{-1, 0}: {0, -1},
	}

	for {
		newPosition := position.Sum(&direction)
		if newPosition.x < 0 || newPosition.y < 0 || newPosition.x >= gridSizeX || newPosition.y >= gridSizeY {
			// guard left the map
			break
		}

		if obstructionsMap[newPosition] {
			// rotate guard
			direction = rotationRules[direction]
			// fmt.Println("Found obstacle, rotate.")
			// debug(position, direction)
		} else {
			// move the guard
			position = newPosition
			visitedPositions[position] = true
			// debug(position, direction)
		}
	}

	return len(visitedPositions)
}

func debug(position vector, direction vector) {
	fmt.Printf("Current position: %v, current direction: %v\n", position, direction)
}
