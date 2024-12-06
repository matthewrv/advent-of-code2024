package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	obstructionsMap, position, direction := readInput("input.txt")
	visited := countVisitedPositions(obstructionsMap, position, direction)
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

func readInput(fileName string) (obstructionsMap [][]bool, position vector, direction vector) {
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
	for scanner.Scan() {
		row := scanner.Text()
		obstructionsMap = append(obstructionsMap, make([]bool, len(row)))
		for col, token := range row {
			switch token {
			case '#':
				obstructionsMap[line][col] = true
			case '^':
				position = vector{col, line}
				// axis are assumed to increase in direction of increasing index of array
				direction = vector{0, -1}
			}
		}

		line++
	}

	return obstructionsMap, position, direction
}

// part 1

func countVisitedPositions(obstructionsMap [][]bool, position vector, direction vector) int {
	visitedPositions := map[vector]bool{}
	visitedPositions[position] = true

	gridSizeX := len(obstructionsMap[0])
	gridSizeY := len(obstructionsMap)

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

		if obstructionsMap[newPosition.y][newPosition.x] {
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
