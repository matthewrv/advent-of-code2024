package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	size, obstructionsMap, start := readInput("input.txt")
	// result := countVisitedPositions(size, obstructionsMap, start)
	result := countCycledRoutes(size, obstructionsMap, start)
	fmt.Printf("Result: %d\n", result)
}

// helper struct

type vector struct {
	x int
	y int
}

func (v1 *vector) Sum(v2 *vector) vector {
	return vector{v1.x + v2.x, v1.y + v2.y}
}

type Visit struct {
	position  vector
	direction vector
}

// read input

func readInput(fileName string) (size vector, obstructionsMap map[vector]bool, start Visit) {
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
				// axis are assumed to increase in direction of increasing index of array
				start = Visit{vector{col, line}, vector{0, -1}}
			}
		}

		size.x = len(row)
		line++
	}

	size.y = line + 1

	return size, obstructionsMap, start
}

// part 1

func countVisitedPositions(size vector, obstructionsMap map[vector]bool, start Visit) int {
	visitedPositions := map[vector]bool{}
	visitedPositions[start.position] = true

	current := start
	for {
		current = getNextPosition(current, obstructionsMap)
		if current.isOutOfMap(size) {
			// guard left the map
			break
		}

		visitedPositions[current.position] = true
	}

	return len(visitedPositions)
}

// part 2

func countCycledRoutes(size vector, obstructionsMap map[vector]bool, start Visit) int {
	possibleObstacles := make(map[vector]bool)

	current := start
	for {
		newPosition := getNextPosition(current, obstructionsMap)
		if newPosition.isOutOfMap(size) {
			// guard left the map
			break
		}

		// debug(newPosition.position, newPosition.direction)
		obstacleExists := obstructionsMap[newPosition.position]
		if !obstacleExists {
			obstructionsMap[newPosition.position] = true
			if isCycle(size, obstructionsMap, start) {
				possibleObstacles[newPosition.position] = true
			}
			obstructionsMap[newPosition.position] = false
		}

		current = newPosition
	}

	return len(possibleObstacles)
}

func isCycle(size vector, obstructionsMap map[vector]bool, start Visit) bool {
	visitedPositions := make(map[Visit]bool)
	visitedPositions[start] = true

	position := start
	for {
		newPosition := getNextPosition(position, obstructionsMap)
		if newPosition.isOutOfMap(size) {
			break
		}

		if visitedPositions[newPosition] {
			return true
		}

		visitedPositions[newPosition] = true
		position = newPosition
	}

	return false
}

func debug(position vector, direction vector) {
	fmt.Printf("Current position: %v, current direction: %v\n", position, direction)
}

func getNextPosition(position Visit, obstructionsMap map[vector]bool) (newVisit Visit) {
	newPosition := position.position.Sum(&position.direction)

	rotationRules := map[vector]vector{
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{0, 1}:  {-1, 0},
		{-1, 0}: {0, -1},
	}

	if obstructionsMap[newPosition] {
		// rotate guard
		newVisit = Visit{position: position.position, direction: rotationRules[position.direction]}
		// fmt.Println("Found obstacle, rotate.")
		// debug(position, direction)
	} else {
		// move the guard
		newVisit = Visit{position: newPosition, direction: position.direction}
		// debug(position, direction)
	}

	return newVisit
}

func (v *Visit) isOutOfMap(size vector) bool {
	if v.position.x < 0 || v.position.y < 0 || v.position.x >= size.x || v.position.y >= size.y {
		// guard left the map
		return true
	}
	return false
}
