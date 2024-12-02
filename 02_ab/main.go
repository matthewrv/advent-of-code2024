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

// change this constant to switch between first and second part
const withDampener bool = true

func main() {
	reports := readFile("input.txt")
	fmt.Printf("Reports read from input file: %d\n", len(reports))
	safeReportsCount := countSafeReports(reports)
	fmt.Println(safeReportsCount)
}

func readFile(fileName string) (reports [][]int) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		var report []int
		for _, token := range tokens {
			number, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, number)
		}
		reports = append(reports, report)
	}

	return reports
}

func countSafeReports(reports [][]int) (count int) {
	for _, report := range reports {
		// fmt.Printf("Processing report %d\n", idx)  // for debug purpose only
		count += isSafeReport(report)
	}
	return count
}

func isSafeReport(report []int) int { // return 1 if safe and 0 if not
	if withDampener {
		return isSafeWithDampener(report)
	}
	isSafe, _ := isSafeDefault(report)
	return isSafe
}

func isSafeDefault(report []int) (int, int) { // return 1 if safe and 0 if not
	var isAscending int = 1
	if report[1]-report[0] < 0 {
		isAscending = -1
	}
	for idx, value := range report[:len(report)-1] {
		diff := (report[idx+1] - value) * isAscending
		if diff > 3 || diff < 1 {
			// fmt.Printf("Failed on values %d, %d\n", value, report[idx+1])  // for debug purpose only
			return 0, idx
		}
	}
	// fmt.Println("Success")  // for debug purpose only
	return 1, 0
}

func isSafeWithDampener(report []int) int { // return 1 if safe and 0 if not
	isSafe, problemIdx := isSafeDefault(report)
	if isSafe == 1 {
		return isSafe
	}

	if problemIdx == 1 {
		isSafe, _ = isSafeDefault(report[1:])
		if isSafe == 1 {
			return isSafe
		}
	}

	testReport1 := slices.Concat(report[:problemIdx], report[problemIdx+1:])
	isSafe, _ = isSafeDefault(testReport1)
	if isSafe == 1 {
		return isSafe
	}

	testReport2 := slices.Concat(report[:problemIdx+1], report[problemIdx+2:])
	isSafe, _ = isSafeDefault(testReport2)
	return isSafe
}
