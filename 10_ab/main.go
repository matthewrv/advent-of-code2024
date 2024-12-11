package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	trailsMap := readInput("input.txt")
	// fmt.Println(trailsMap)

	// score := getScore(trailsMap)
	// fmt.Println("Result: ", score)

	rating := getRating(trailsMap)
	fmt.Println("Result: ", rating)
}

func readInput(fileName string) (trailsMap Map) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	trailsMap.Map = [][]int{}
	height := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var (
			line      = scanner.Text()
			tokens    = strings.Split(line, "")
			trailsRow []int
		)
		for _, token := range tokens {
			val, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			trailsRow = append(trailsRow, val)
		}

		trailsMap.Map = append(trailsMap.Map, trailsRow)
		trailsMap.width = len(trailsRow)
		height++
	}
	trailsMap.height = height

	return trailsMap
}

type Map struct {
	Map    [][]int
	width  int
	height int
}

// part 1

func getScore(trailsMap Map) (score int) {
	for i, row := range trailsMap.Map {
		for j, point := range row {
			if point == 0 {
				score += len(findTops(trailsMap, i, j))
			}
		}
	}

	return score
}

var directions = [4][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func findTops(trailsMap Map, row int, col int) (tops map[[2]int]bool) {
	current := trailsMap.Map[row][col]

	if current == 9 {
		return map[[2]int]bool{
			{row, col}: true,
		}
	}

	tops = map[[2]int]bool{}
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := col + direction[1]

		if nextRow < 0 || nextCol < 0 || nextRow >= trailsMap.height || nextCol >= trailsMap.width {
			continue
		}

		if trailsMap.Map[nextRow][nextCol] == current+1 {
			for top := range findTops(trailsMap, nextRow, nextCol) {
				tops[top] = true
			}
		}
	}

	return tops
}

// part 2

func getRating(trailsMap Map) (score int) {
	for i, row := range trailsMap.Map {
		for j, point := range row {
			if point == 0 {
				score += findRoutes(trailsMap, i, j)
			}
		}
	}

	return score
}

func findRoutes(trailsMap Map, row int, col int) (routes int) {
	current := trailsMap.Map[row][col]

	if current == 9 {
		return 1
	}

	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := col + direction[1]

		if nextRow < 0 || nextCol < 0 || nextRow >= trailsMap.height || nextCol >= trailsMap.width {
			continue
		}

		if trailsMap.Map[nextRow][nextCol] == current+1 {
			routes += findRoutes(trailsMap, nextRow, nextCol)
		}
	}

	return routes
}
