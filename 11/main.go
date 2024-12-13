package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	stones := readInput("input.txt")
	fmt.Println(stones)
	iters := 75
	for range iters {
		stones = blink(stones)
		// fmt.Println("Finished iteration:", iter, "Current length:", len(stones))
		// fmt.Println(input)
	}

	var totalStones int
	for _, count := range stones {
		totalStones += count
	}
	fmt.Println("Result: ", totalStones)
}

func readInput(fileName string) (stones map[int]int) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	stones = make(map[int]int)
	for _, token := range bytes.Split(content, []byte(" ")) {
		val, err := strconv.Atoi(string(token))
		if err != nil {
			log.Fatal(err)
		}
		stones[val] += 1
	}

	return stones
}

func blink(stones map[int]int) (newStones map[int]int) {
	cache := map[int][]int{}

	newStones = make(map[int]int)
	for value, count := range stones {
		transformSingleValue(newStones, value, count, cache)
	}

	return newStones
}

func transformSingleValue(newStones map[int]int, value int, count int, cache map[int][]int) {
	if cache[value] == nil {
		var tmpStones []int

		str := fmt.Sprintf("%d", value)
		switch {
		case value == 0:
			tmpStones = []int{1}
		case len(str)%2 == 0:
			val1, _ := strconv.Atoi(str[:len(str)/2])
			val2, _ := strconv.Atoi(str[len(str)/2:])
			tmpStones = []int{val1, val2}
		default:
			tmpStones = []int{value * 2024}
		}

		cache[value] = tmpStones
	}

	for _, newStone := range cache[value] {
		newStones[newStone] += count
	}
}
