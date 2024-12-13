package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

// switch between part 1 and part 2
const part1 bool = false

func main() {
	gardenMap := readInput("input.txt")
	cost := getCost(gardenMap)
	fmt.Println("Result: ", cost)
}

// helper structs

type GardenMap struct {
	plots     [][]string
	totalRows int
	totalCols int
}

type Region struct {
	value string
	area  int
	plots [][2]int
	edges map[int]bool // external edges of region
}

// reading input

func readInput(fileName string) (gardenMap GardenMap) {
	gardenMap.plots = make([][]string, 0)

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
		gardenMap.plots = append(gardenMap.plots, strings.Split(input.Text(), ""))
	}

	gardenMap.totalRows = len(gardenMap.plots)
	gardenMap.totalCols = len(gardenMap.plots[0])

	return gardenMap
}

// solution

func getCost(gardenMap GardenMap) (cost int) {

	regions := buildRegions(gardenMap)
	for _, region := range regions {
		cost += getFenceCost(region, gardenMap)
	}

	return cost
}

func getFenceCost(region *Region, gardenMap GardenMap) (cost int) {
	// debug info
	if part1 {
		cost = region.area * len(region.edges)
	} else {
		sorted := slices.Sorted(maps.Keys(region.edges))
		prev := sorted[0]
		totalEdges := 1
		for _, edge := range sorted[1:] {
			if edge-prev != 1 || (prev+1)%gardenMap.totalRows == 0 || splitByAnotherEdge(prev, gardenMap, region) { // assume square gardenMap
				totalEdges++
			}
			prev = edge
		}
		cost = region.area * totalEdges
	}
	return cost
}

func splitByAnotherEdge(edge int, gardenMap GardenMap, region *Region) bool {
	for _, neighbour := range findNeighbourEdges(edge, gardenMap) {
		if region.edges[neighbour] {
			return true
		}
	}

	return false
}

func buildRegions(gardenMap GardenMap) (regions []*Region) {
	plotToRegionMap := make(map[[2]int]*Region)

	collideRegions := func(reg1 *Region, reg2 *Region) {
		// update parameters from one region to other
		reg1.area += reg2.area
		reg1.plots = append(reg1.plots, reg2.plots...)
		maps.Copy(reg1.edges, reg2.edges)

		// pop one region from list of regions
		idx := slices.Index(regions, reg2)
		regions = slices.Delete(regions, idx, idx+1)

		// replace all references to removed region
		for _, plot := range reg2.plots {
			plotToRegionMap[plot] = reg1
		}
	}

	for i, row := range gardenMap.plots {
		for j, val := range row {
			var (
				region *Region = nil
				plot           = [2]int{i, j}
				edges  []int   = plotEdges(i, j, gardenMap)
			)

			if i > 0 && gardenMap.plots[i-1][j] == val {
				prev := [2]int{i - 1, j}
				region = plotToRegionMap[prev]
			}
			if j > 0 && gardenMap.plots[i][j-1] == val {
				prev := [2]int{i, j - 1}
				newRegion := plotToRegionMap[prev]

				if region == nil {
					region = newRegion
				} else if region != newRegion {
					collideRegions(region, newRegion)
				}
			}
			if region == nil {
				region = new(Region)
				region.value = val
				region.edges = make(map[int]bool)
				regions = append(regions, region)
			}

			region.area++
			region.plots = append(region.plots, plot)

			for _, edge := range edges {
				if region.edges[edge] {
					// common edge - internal. Do not include in perimeter
					delete(region.edges, edge)
				} else {
					region.edges[edge] = true
				}
			}

			plotToRegionMap[plot] = region
		}
	}

	return regions
}

// All the magic of solution - enumerate all edges of plots
//
// For horizontal lines use this order
//
//   1  2  3  4
//   v  v  v  v
//  __ __ __ __
// |  |  |  |  |
// | 5| 6| 7| 8|
//  __ __ __ __
//
// For vertical lines we continue numbering from last horizontal edge
//
//        __ __ __ __
// 9  -> |10|11|12|  | <- 13
//       |  |  |  |  |
//        __ __ __ __
//
// So for grid 1x4 plots (rows x cols) nodes, we total have (1+1)*4=8 horizontal edges and 1*(4+1)=5 vertical edges.
// And in total 13 edges.

func plotEdges(i int, j int, gardenMap GardenMap) []int {
	m, n := gardenMap.totalRows, gardenMap.totalCols
	allHorizontalEdges := n * (m + 1)

	return []int{
		i*n + j,
		(i+1)*n + j,
		allHorizontalEdges + j*m + i,
		allHorizontalEdges + (j+1)*m + i,
	}
}

func findNeighbourEdges(edge int, gardenMap GardenMap) (edges []int) {
	m, n := gardenMap.totalRows, gardenMap.totalCols
	allHorizontalEdges := n * (m + 1)

	isHorizontal := edge < allHorizontalEdges
	if isHorizontal {
		i := edge / n
		j := edge % n

		if i > 0 {
			edges = append(edges, allHorizontalEdges+(j+1)*m+i-1)
		}
		if i < gardenMap.totalRows {
			edges = append(edges, allHorizontalEdges+(j+1)*m+i)
		}
	} else {
		i := (edge - allHorizontalEdges) % m
		j := (edge - allHorizontalEdges) / m

		if j > 0 {
			edges = append(edges, (i+1)*n+j-1)
		}
		if j < gardenMap.totalCols {
			edges = append(edges, (i+1)*n+j)
		}
	}

	return edges
}
