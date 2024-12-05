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
	rules, updates := readInput("input.txt")
	// fmt.Printf("%v\n%v\n", rules, updates)
	result := findValidUpdates(rules, updates)
	fmt.Printf("Result: %d\n", result)
}

// helper struct

type Set map[int]bool

func add(set Set, value int) {
	set[value] = true
}

func in(set Set, value int) bool {
	return set[value]
}

// read input

func readInput(fileName string) (rules map[int]Set, updates [][]int) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rules = map[int]Set{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() && len(scanner.Text()) != 0 {
		var (
			line   = scanner.Text()
			tokens = strings.Split(line, "|")
			a, _   = strconv.Atoi(tokens[0])
			b, _   = strconv.Atoi(tokens[1])
		)
		if rules[b] == nil {
			rules[b] = Set{}
		}
		add(rules[b], a)
	}

	for scanner.Scan() {
		var (
			line   = scanner.Text()
			tokens = strings.Split(line, ",")
			update []int
		)
		for _, token := range tokens {
			page, _ := strconv.Atoi(token)
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

// find valid updates according to rules and return sum of middle elements
func findValidUpdates(rules map[int]Set, updates [][]int) (total int) {
	for _, update := range updates {
		total += checkOneUpdate(rules, update)
	}
	return total
}

func checkOneUpdate(rules map[int]Set, update []int) int {
	for idx, page := range update {
		for i := 0; i < idx; i++ {
			if in(rules[update[i]], page) {
				return 0
			}
		}
	}

	mid := len(update) / 2
	fmt.Printf("Middle %d, array %v\n", update[mid], update)
	return update[mid]
}
