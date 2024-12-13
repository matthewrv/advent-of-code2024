package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	arcades := readInput("./input.txt")
	total := winEverything(arcades)
	fmt.Println("Result:", total)
}

type Arcade struct {
	ax int
	ay int
	bx int
	by int
	x  int
	y  int
}

func readInput(fileName string) (arcades []Arcade) {
	reButton := regexp.MustCompile("Button [AB]: X\\+([0-9]+), Y\\+([0-9]+)")
	rePrize := regexp.MustCompile("Prize: X=([0-9]+), Y=([0-9]+)")

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	i := 0
	for i < len(lines) {
		var arcade Arcade
		arcade.ax, arcade.ay = parseCondition(reButton, lines[i])
		arcade.bx, arcade.by = parseCondition(reButton, lines[i+1])
		arcade.x, arcade.y = parseCondition(rePrize, lines[i+2])

		// for part 2 only
		arcade.x += 10000000000000
		arcade.y += 10000000000000

		i += 4

		arcades = append(arcades, arcade)
	}

	return arcades
}

func parseCondition(re *regexp.Regexp, str string) (x int, y int) {
	match := re.FindStringSubmatch(str)
	x, _ = strconv.Atoi(match[1])
	y, _ = strconv.Atoi(match[2])
	return x, y
}

func winEverything(arcades []Arcade) (result int) {
	for _, arcade := range arcades {
		result += solveArcade(&arcade)
	}
	return result
}

func solveArcade(arcade *Arcade) (tokensToSpent int) {
	fmt.Printf("Solving arcade: %+v\n", arcade)

	det := arcade.ax*arcade.by - arcade.ay*arcade.bx
	if det == 0 {
		return 0
	}

	aDet := arcade.by*arcade.x - arcade.bx*arcade.y
	if aDet%det != 0 {
		return 0
	}

	bDet := -arcade.ay*arcade.x + arcade.ax*arcade.y
	if bDet%det != 0 {
		return 0
	}

	aTokens := aDet / det
	bTokens := bDet / det

	// for part 1 only
	// if aTokens > 100 || bTokens > 100 {
	// 	return 0
	// }

	fmt.Printf("Solution: %d, %d\n", aTokens, bTokens)

	return 3*aTokens + bTokens
}
