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
	result := solvePart1("input.txt")
	fmt.Printf("Result: %v\n", result)
}

// part 1

func solvePart1(fileName string) (result int) {
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
	idx := 0
	for scanner.Scan() {
		var (
			line     = scanner.Text()
			tmp      = strings.Split(line, ": ")
			total, _ = strconv.Atoi(tmp[0])
			tokens   = strings.Split(tmp[1], " ")
		)
		var (
			elements []int
		)
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			elements = append(elements, num)
		}
		idx++

		if bruteForce(total, elements[0], elements[1:]) {
			result += total
		}
	}

	return result
}

func bruteForce(total int, current int, elements []int) bool {
	if len(elements) == 1 {
		if current+elements[0] == total || current*elements[0] == total {
			return true
		}
		return false
	}

	return bruteForce(total, current+elements[0], elements[1:]) || bruteForce(total, current*elements[0], elements[1:])
}
