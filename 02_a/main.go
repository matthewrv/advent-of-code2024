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
		count += isSafeReport(report)
	}
	return count
}

func isSafeReport(report []int) int { // return 1 if safe and 0 if not
	var isAscending int = 1
	if report[1]-report[0] < 0 {
		isAscending = -1
	}
	for idx, value := range report[:len(report)-1] {
		diff := (report[idx+1] - value) * isAscending
		if diff > 3 || diff < 1 {
			return 0
		}
	}
	return 1
}
