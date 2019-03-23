package main

import (
	"fmt"
	"os"
	regex "regexp"
	"strconv"
	"strings"
)

var (
	input = strings.Split(os.Getenv("ADVENT_INPUT"), "\n")
)

func main() {
	fmt.Println(findDupSquares(input))
}

func findDupSquares(input []string) int {
	seen := make(map[string]int)
	dups := 0

	for _, claim := range input {
		claimVals := regex.MustCompile("\\s@\\s|:\\s|x|,").Split(claim, -1)

		startTop, _ := strconv.Atoi(claimVals[2])
		startLeft, _ := strconv.Atoi(claimVals[1])
		width, _ := strconv.Atoi(claimVals[3])
		height, _ := strconv.Atoi(claimVals[4])

		for i := startTop; i < startTop+height; i++ {
			for j := startLeft; j < startLeft+width; j++ {
				seen[fmt.Sprintf("%v,%v", i, j)]++
			}
		}
	}

	for _, v := range seen {
		if v > 1 {
			dups++
		}
	}

	return dups
}
