package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	input := readInput("input.txt")
	fmt.Println(input)
	iters := 25
	for iter := range iters {
		input = transform(input)
		fmt.Println("Finished iteration:", iter, "Current length:", len(input))
		// fmt.Println(input)
	}

	fmt.Println("Result: ", len(input))
}

func readInput(fileName string) (arr []int) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, token := range bytes.Split(content, []byte(" ")) {
		val, err := strconv.Atoi(string(token))
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, val)
	}

	return arr
}

func transform(input []int) (newArr []int) {
	cache := map[int][]int{}

	subSlices := [][]int{}
	for _, value := range input {
		subSlices = append(subSlices, transformSingleValue(value, cache))
	}

	return slices.Concat(subSlices...)
}

func transformSingleValue(value int, cache map[int][]int) (result []int) {
	if cache[value] != nil {
		return cache[value]
	}

	str := fmt.Sprintf("%d", value)
	switch {
	case value == 0:
		result = []int{1}
	case len(str)%2 == 0:
		val1, _ := strconv.Atoi(str[:len(str)/2])
		val2, _ := strconv.Atoi(str[len(str)/2:])
		result = []int{val1, val2}
	default:
		result = []int{value * 2024}
	}

	cache[value] = result
	return result
}
