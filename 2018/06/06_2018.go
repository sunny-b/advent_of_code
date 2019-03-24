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
	fmt.Println(findUniqueClaim(input))
}

func findUniqueClaim(input []string) string {
	seen := make(map[string]int)

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

	for _, claim := range input {
		overlaps := false
		claimVals := regex.MustCompile("\\s@\\s|:\\s|x|,").Split(claim, -1)

		startTop, _ := strconv.Atoi(claimVals[2])
		startLeft, _ := strconv.Atoi(claimVals[1])
		width, _ := strconv.Atoi(claimVals[3])
		height, _ := strconv.Atoi(claimVals[4])

		for i := startTop; i < startTop+height; i++ {
			for j := startLeft; j < startLeft+width; j++ {
				if seen[fmt.Sprintf("%v,%v", i, j)] > 1 {
					overlaps = true
				}
			}
		}

		if !overlaps {
			return claimVals[0]
		}
	}

	return ""
}
