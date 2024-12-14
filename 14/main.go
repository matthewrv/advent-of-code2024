package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	positions, velocities, gridSize := readInput("input.txt")
	// fmt.Printf("Positions: %v, velocities %v\n", positions, velocities)
	results := calcSafetyFactor(positions, velocities, gridSize)
	fmt.Println("Result:", results)
}

func readInput(fileName string) (positions [][2]int, velocities [][2]int, gridSize [2]int) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	re := regexp.MustCompile("p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)")
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		match := re.FindStringSubmatch(line)
		var (
			x, _  = strconv.Atoi(match[1])
			y, _  = strconv.Atoi(match[2])
			vx, _ = strconv.Atoi(match[3])
			vy, _ = strconv.Atoi(match[4])
		)
		positions = append(positions, [2]int{x, y})
		velocities = append(velocities, [2]int{vx, vy})
	}
	gridSize = [2]int{101, 103}

	return positions, velocities, gridSize
}

func calcSafetyFactor(positions [][2]int, velocities [][2]int, gridSize [2]int) (safetyFactor int) {
	var newPositions [][2]int = make([][2]int, len(positions))
	for sec := range 9999 {
		for i, position := range positions {
			newPositions[i] = updatePosition(position, velocities[i], gridSize)
		}
		newPositions, positions = positions, newPositions

		printPositions(sec+1, positions, gridSize)
	}

	// split
	regions := [4]int{}
	middleX := gridSize[0] / 2
	middleY := gridSize[1] / 2
	for _, position := range positions {
		switch {
		case position[0] < middleX && position[1] < middleY:
			regions[0]++
		case position[0] < middleX && position[1] > middleY:
			regions[1]++
		case position[0] > middleX && position[1] < middleY:
			regions[2]++
		case position[0] > middleX && position[1] > middleY:
			regions[3]++
		}
	}

	// fmt.Println(regions)
	// fmt.Println(positions)

	safetyFactor = regions[0]
	for _, val := range regions[1:] {
		safetyFactor *= val
	}

	return safetyFactor
}

func updatePosition(position [2]int, velocity [2]int, gridSize [2]int) [2]int {
	x := (position[0] + velocity[0]) % gridSize[0]
	if x < 0 {
		x = gridSize[0] + x
	}
	y := (position[1] + velocity[1]) % gridSize[1]
	if y < 0 {
		y = gridSize[1] + y
	}
	return [2]int{x, y}
}

func printPositions(iter int, positions [][2]int, gridSize [2]int) {
	rect := image.Rect(0, 0, gridSize[0], gridSize[1])
	palette := []color.Color{color.Black, color.RGBA{0, 144, 0, 255}}
	img := image.NewPaletted(rect, palette)
	for _, position := range positions {
		img.SetColorIndex(position[0], position[1], 1)
	}

	f, err := os.Create(fmt.Sprintf("images/%04d.png", iter))
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, img)
	f.Close()
}
