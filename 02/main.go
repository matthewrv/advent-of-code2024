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

// Struct for testing different implementations

type Solver int

const (
	V1Default Solver = iota
	V1WithDampener
	V2Default
	V2WithDampener
)

var solverName = map[Solver]string{
	V1Default:      "v1 default",
	V1WithDampener: "v1 with dampener",
	V2Default:      "v2 default",
	V2WithDampener: "v2 with dampener",
}

func (ss Solver) String() string {
	return solverName[ss]
}

// Solution

const selectedSolver = V2WithDampener

func main() {
	reports := readFile("input.txt")
	fmt.Printf("Reports read from input file: %d\n", len(reports))

	safeReportsCount := countSafeReports(reports)
	fmt.Printf("Solution %v. Found safe reports: %v\n", selectedSolver, safeReportsCount)
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

func isSafeReport(report []int) (isSafe int) { // return 1 if safe and 0 if not
	switch selectedSolver {
	case V1Default:
		isSafe, _ := isSafeDefault(report)
		return isSafe
	case V1WithDampener:
		return isSafeWithDampener(report)
	case V2Default:
		diff := getFirstDiff(report)
		isSafe, _ := isSafeDefaultV2(diff)
		return isSafe
	case V2WithDampener:
		diff := getFirstDiff(report)
		return isSafeWithDampenerV2(diff)
	}
	return isSafe
}

// First implementation - direct approach

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

// Second impelentation - caclulate diffs only once for each report

func isSafeWithDampenerV2(diff []int) int { // return 1 if safe and 0 if not
	isSafe, problemIdx := isSafeDefaultV2(diff)
	if isSafe == 1 {
		return isSafe
	}

	// can only mean that absolute value of difference between two first values is too big
	// (ascending/descending condition could not be failed for first diff)
	if problemIdx == 0 {
		isSafe, _ = isSafeDefaultV2(diff[1:])
		return isSafe
	}

	if problemIdx == len(diff)-1 {
		// we can just ignore one last element
		return 1
	}

	testdiff1 := reduceDiff(diff, problemIdx-1)
	isSafe, _ = isSafeDefaultV2(testdiff1)
	if isSafe == 1 {
		return isSafe
	}

	testdiff2 := reduceDiff(diff, problemIdx)
	isSafe, _ = isSafeDefaultV2(testdiff2)
	return isSafe
}

func isSafeDefaultV2(diff []int) (int, int) {
	var isAscending int = 1
	if diff[0] < 0 {
		isAscending = -1
	}

	for idx, value := range diff {
		current_diff := isAscending * value
		if current_diff > 3 || current_diff < 1 {
			// fmt.Printf("Failed on values %d, %d\n", value, report[idx+1])  // for debug purpose only
			return 0, idx
		}
	}
	return 1, 0
}

func getFirstDiff(report []int) (diff []int) {
	for i := 0; i < len(report)-1; i++ {
		diff = append(diff, report[i+1]-report[i])
	}
	return diff
}

// Sum element at index `at` with next one and return newly created slice.
//
// Semantically equivalent to removing one element from original array and calculating new diff
func reduceDiff(diff []int, at int) []int {
	var leading []int
	if at > 0 {
		leading = diff[:at-1]
	}
	return slices.Concat(leading, []int{diff[at] + diff[at+1]}, diff[at+2:])
}
