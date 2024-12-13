package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const part int = 2 // switch between part 1 and 2 of puzzle

func main() {
	re_do := regexp.MustCompile("do\\(\\)")
	re_dont := regexp.MustCompile("don't\\(\\)")
	re_mul := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var parsed [][][]byte
	total := 0
	switch part {
	case 1:
		parsed = re_mul.FindAllSubmatch(content, -1)
		total = getSum(parsed)
	case 2:
		next_re := re_dont
		current_re := re_do
		from := 0

		for match := next_re.FindIndex(content[from:]); match != nil; match = next_re.FindIndex(content[from:]) {
			// fmt.Printf("%v\n", current_re == re_do) // for debug purpose
			if current_re == re_do {
				parsed = re_mul.FindAllSubmatch(content[from:from+match[0]], -1)
				total += getSum(parsed)
			}

			from += match[1]
			current_re, next_re = next_re, current_re
			// fmt.Printf("equal: %v\n", current_re == next_re) // for debug purpose
		}

		// process tail
		if current_re == re_do {
			parsed = re_mul.FindAllSubmatch(content[from:], -1)
			total += getSum(parsed)
		}
	}

	fmt.Printf("Result: %d\n", total)
}

func getSum(parsed [][][]byte) int {
	total := 0
	for _, match := range parsed {
		var (
			a, _ = strconv.Atoi(string(match[1]))
			b, _ = strconv.Atoi(string(match[2]))
		)
		// fmt.Printf("%d * %d\n", a, b) // for debug purpose
		total += a * b
	}

	return total
}
