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
			maxValue int
		)
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			elements = append(elements, num)
			if maxValue == 0 {
				maxValue += num
			} else if num == 1 {
				maxValue += 1
			} else {
				maxValue *= num
			}
		}
		idx++

		// solve line

		// edge cases
		if maxValue < total {
			fmt.Printf("Non-solvable line %d: %d - %v | %d\n", idx, total, elements, maxValue)
			continue
		}
		if maxValue == total {
			result += total
			continue
		}

		// general case
		// stack := []Operation{}
		// for _ = range elements {
		// 	stack = append(stack, Product)
		// }
		// stack = stack[:len(stack)-1]
		result += generalCase(total, elements[0], elements[1:])

	}

	return result
}

type Operation int

const (
	Sum Operation = iota
	Product
)

func generalCase(total int, current int, elements []int) (closest int) {
	val := elements[0]

	if len(elements) == 1 {
		if current*val == total {
			return total
		}
		if current+val == total {
			return total
		}
		return 0
	}

	closest = generalCase(total, current*val, elements[1:])
	if closest == total {
		return total
	}
	closest = generalCase(total, current+val, elements[1:])
	if closest == total {
		return total
	}

	return 0
}

// func getMaxResult(current int, element int) int {
// 	if element == 1 {
// 		return current + element
// 	}
// 	return current * element
// }

// func getMinResult(current int, element int) int {
// 	if element == 1 {
// 		return current * element
// 	}
// 	return current + element
// }
