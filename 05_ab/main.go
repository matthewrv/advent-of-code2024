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
	order, updates := readInput("input.txt")
	// fmt.Printf("%v\n%+v\n", updates, order)
	// result := findValidUpdates(order, updates)
	result := fixInvalidUpdates(order, updates)
	fmt.Printf("Result: %d\n", result)
}

// helper struct

type PagePair struct {
	page1 int
	page2 int
}

type PageOrdering struct {
	order map[PagePair]int
}

func (order *PageOrdering) comparePages(page1 int, page2 int) int {
	return order.order[PagePair{page1, page2}]
}

// helper func

func getUpdateMid(update []int) int {
	mid := len(update) / 2
	fmt.Printf("Middle %d, array %v\n", update[mid], update)
	return update[mid]
}

// read input

func readInput(fileName string) (pageOrder PageOrdering, updates [][]int) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	pageOrder.order = map[PagePair]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() && len(scanner.Text()) != 0 {
		var (
			line   = scanner.Text()
			tokens = strings.Split(line, "|")
			a, _   = strconv.Atoi(tokens[0])
			b, _   = strconv.Atoi(tokens[1])
		)

		pageOrder.order[PagePair{a, b}] = -1
		pageOrder.order[PagePair{b, a}] = 1
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

	return pageOrder, updates
}

// Part 1

// find valid updates according to ordering and return sum of middle elements
func findValidUpdates(pageOrder PageOrdering, updates [][]int) (total int) {
	for _, update := range updates {
		total += checkOneUpdate(pageOrder, update)
	}
	return total
}

// check if update is correct and if so - return middle element, otherwise return 0
func checkOneUpdate(pageOrder PageOrdering, update []int) int {
	if slices.IsSortedFunc(update, pageOrder.comparePages) {
		return getUpdateMid(update)
	}

	return 0
}

// Part 2

func fixInvalidUpdates(pageOrder PageOrdering, updates [][]int) (total int) {
	for _, update := range updates {
		if checkOneUpdate(pageOrder, update) == 0 {
			total += fixOneUpdate(pageOrder, update)
		}
	}
	return total
}

// fix sorting of update and return value of middle element
func fixOneUpdate(pageOrder PageOrdering, update []int) int {
	slices.SortFunc(update, pageOrder.comparePages)
	return getUpdateMid(update)
}
