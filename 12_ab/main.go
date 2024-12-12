package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	gardenMap := readInput("input.txt")
	cost := getCost(gardenMap)
	fmt.Println("Result: ", cost)
}

func readInput(fileName string) (gardenMap [][]string) {
	gardenMap = make([][]string, 0)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	input := bufio.NewScanner(f)
	for input.Scan() {
		gardenMap = append(gardenMap, strings.Split(input.Text(), ""))
	}

	return gardenMap
}

type Region struct {
	value     string
	area      int
	perimeter int
	nodes     [][2]int
}

func getCost(gardenMap [][]string) (cost int) {
	regions := []*Region{}
	checked := make(map[[2]int]*Region)

	collideRegions := func(reg1 *Region, reg2 *Region) {
		// update parameters
		reg1.area += reg2.area
		reg1.perimeter += reg2.perimeter
		reg1.nodes = append(reg1.nodes, reg2.nodes...)

		// pop one from all regions
		idx := slices.Index(regions, reg2)
		regions = slices.Delete(regions, idx, idx+1)

		// replace all references to removed region
		for _, node := range reg2.nodes {
			checked[node] = reg1
		}
	}

	for i, row := range gardenMap {
		for j, val := range row {
			var (
				region      *Region
				commonSides int
			)

			if i > 0 && gardenMap[i-1][j] == val {
				prev := [2]int{i - 1, j}
				region = checked[prev]
				commonSides++
			}
			if j > 0 && gardenMap[i][j-1] == val {
				prev := [2]int{i, j - 1}
				newRegion := checked[prev]

				if region == nil {
					region = newRegion
				} else if region != newRegion {
					collideRegions(region, newRegion)
				}
				commonSides++
			}
			if region == nil {
				region = new(Region)
				region.value = val
				regions = append(regions, region)
			}

			region.area++
			region.perimeter = region.perimeter - commonSides + (4 - commonSides)

			position := [2]int{i, j}
			checked[position] = region
			region.nodes = append(region.nodes, position)
		}
	}

	for _, region := range regions {
		cost += getFenceCost(region)
	}

	return cost
}

func getFenceCost(region *Region) (cost int) {
	// debug info
	// fmt.Printf("Region %q with area %d and perimeter %d: %v\n", region.value, region.area, region.perimeter, region.nodes)
	return region.area * region.perimeter
}
