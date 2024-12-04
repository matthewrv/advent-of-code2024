package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	newSlice := slice[:2]
	newSlice = append(newSlice, 5)
	fmt.Printf("Original: %+v\n", slice)
	fmt.Printf("New: %+v\n", newSlice)
}
