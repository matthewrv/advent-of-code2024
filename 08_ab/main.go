package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	nodesLocation, mapSize := readInput("input_test.txt")
	antinodesCount := countAntinodes(nodesLocation, mapSize)
	fmt.Printf("Result: %d\n", antinodesCount)
}

type location struct {
	i int
	j int
}

func (loc1 *location) Sub(loc2 *location) location {
	return location{loc1.i - loc2.i, loc1.j - loc2.j}
}

func (loc1 *location) Sum(loc2 *location) location {
	return location{loc1.i + loc2.i, loc1.j + loc2.j}
}

func readInput(fileName string) (result map[byte][]location, size [2]int) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	result = map[byte][]location{}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		size[1] = len(line)
		for j, sym := range line {
			if sym != '.' {
				loc := location{i, j}
				result[sym] = append(result[sym], loc)
			}
		}

		i++
	}
	size[0] = i

	fmt.Printf("locations: %+v\n", result)

	return result, size
}

func countAntinodes(result map[byte][]location, mapSize [2]int) int {
	uniqueAntinodes := map[location]bool{}
	for _, locations := range result {
		for idx, loc1 := range locations {
			for _, loc2 := range locations[idx+1:] {
				antinodes := getAntinodesOf2(loc1, loc2, mapSize)
				for _, antinode := range antinodes {
					uniqueAntinodes[antinode] = true
				}
			}
		}
	}

	// fmt.Printf("Unique antinodes: %v\n", uniqueAntinodes)
	return len(uniqueAntinodes)
}

func getAntinodesOf2(loc1 location, loc2 location, mapSize [2]int) (result []location) {
	diff := loc2.Sub(&loc1)

	antinode1 := loc1.Sub(&diff)
	if antinode1.inMap(&mapSize) {
		result = append(result, antinode1)
	}

	antinode2 := loc2.Sum(&diff)
	if antinode2.inMap(&mapSize) {
		result = append(result, antinode2)
	}

	return result
}

func (loc *location) inMap(mapSize *[2]int) bool {
	return loc.i >= 0 && loc.j >= 0 && loc.i < mapSize[0] && loc.j < mapSize[1]
}
