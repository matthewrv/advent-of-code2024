package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile("mul\\((?P<first>[0-9]{1,3}),(?P<second>[0-9]{1,3})\\)")
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parsed := re.FindAllSubmatch(content, -1)

	total := 0
	for _, match := range parsed {
		var (
			a, _ = strconv.Atoi(string(match[1]))
			b, _ = strconv.Atoi(string(match[2]))
		)
		total += a * b
	}
	fmt.Printf("%d\n", total)
}
