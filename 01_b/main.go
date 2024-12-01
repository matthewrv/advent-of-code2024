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
	arr1, counter2 := readFile("input.txt")
	similarity := calcSimilarity(arr1, counter2)
	fmt.Println(similarity)
}

func readFile(fileName string) ([]int, map[int]int) {
	f, err := os.Open(fileName)
	check(err)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		scanner  = bufio.NewScanner(f)
		arr1     [1000]int
		size     int
		counter2 = map[int]int{}
	)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "   ")

		a, err := strconv.Atoi(tokens[0])
		check(err)
		arr1[size] = a
		size += 1

		b, err := strconv.Atoi(tokens[1])
		check(err)
		counter2[b] = counter2[b] + 1
	}
	return arr1[:size], counter2
}

func calcSimilarity(a []int, b map[int]int) (similarity int) {
	for _, value := range a {
		similarity += value * b[value]
	}
	return similarity
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
