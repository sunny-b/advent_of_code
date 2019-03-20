package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	input = strings.Split(os.Getenv("ADVENT_INPUT"), "\n")
)

func main() {
	fmt.Println(generateChecksum(input))
}

func generateChecksum(input []string) int {
	twos := 0
	threes := 0

	for _, boxID := range input {
		var (
			twosExist   = false
			threesExist = false
			seen        = make(map[rune]int)
		)

		for _, char := range boxID {
			seen[char]++
		}

		for _, v := range seen {
			if v == 2 {
				twosExist = true
			} else if v == 3 {
				threesExist = true
			}
		}

		if twosExist {
			twos++
		}
		if threesExist {
			threes++
		}
	}

	return twos * threes
}
