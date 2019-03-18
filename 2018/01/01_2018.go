package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	input = strings.Split(os.Getenv("ADVENT_INPUT"), "\n")
)

func main() {
	fmt.Println(updateFrequency(input))
}

func updateFrequency(input []string) int {
	frequency := 0

	for _, update := range input {
		sign := string(update[0])
		num, _ := strconv.Atoi(string(update[1:]))

		switch sign {
		case "+":
			frequency += num
		case "-":
			frequency -= num
		}
	}

	return frequency
}
