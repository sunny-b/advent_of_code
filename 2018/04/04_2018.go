package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	input = strings.Split(os.Getenv("ADVENT_INPUT"), "\n")
)

func main() {
	fmt.Println(findCommonIDs(input))
}

func findCommonIDs(input []string) string {
	for idx, boxID := range input {
		for i := idx; i < len(input); i++ {
			comparison := input[i]
			differences := 0
			differenceLoco := 0

			for i := 0; i < len(boxID); i++ {
				if boxID[i] != comparison[i] {
					differences++
					differenceLoco = i
				}
			}

			if differences == 1 {
				var buffer bytes.Buffer

				buffer.WriteString(boxID[:differenceLoco])
				buffer.WriteString(boxID[differenceLoco+1:])

				return buffer.String()
			}
		}
	}

	return ""
}
