package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	nodesLocation, mapSize := readInput("input.txt")
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
				antinodes := getAntinodesOf2V2(loc1, loc2, mapSize)
				for _, antinode := range antinodes {
					uniqueAntinodes[antinode] = true
				}
			}
		}
	}

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

func getAntinodesOf2V2(loc1 location, loc2 location, mapSize [2]int) (result []location) {
	diff := loc2.Sub(&loc1)
	gcd := GCDEuclidean(diff.i, diff.j)
	step := location{diff.i / gcd, diff.j / gcd}
	// fmt.Printf("%v and %v: diff %v, GCD %d, step %v\n", loc1, loc2, diff, gcd, step)

	current := loc1
	for current.inMap(&mapSize) {
		result = append(result, current)
		current = current.Sub(&step)
	}

	current = loc1.Sum(&step)
	for current.inMap(&mapSize) {
		result = append(result, current)
		current = current.Sum(&step)
	}

	// fmt.Printf("Antinodes are: %v\n", result)

	return result
}

func (loc *location) inMap(mapSize *[2]int) bool {
	return loc.i >= 0 && loc.j >= 0 && loc.i < mapSize[0] && loc.j < mapSize[1]
}

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
