package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	arr1, arr2 := readFile("input.txt")
	slices.Sort(arr1)
	slices.Sort(arr2)
	distance := calcDistance(arr1, arr2)
	fmt.Println(distance)
}

func readFile(fileName string) (arr1 []int, arr2 []int) {
	f, err := os.Open(fileName)
	check(err)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "   ")

		a, err := strconv.Atoi(tokens[0])
		check(err)
		arr1 = append(arr1, a)

		b, err := strconv.Atoi(tokens[1])
		check(err)
		arr2 = append(arr2, b)
	}
	return arr1, arr2
}

func calcDistance(a []int, b []int) (distance int) {
	for idx, value := range a {
		diff := b[idx] - value
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}
	return distance
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
