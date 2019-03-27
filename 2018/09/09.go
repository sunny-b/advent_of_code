package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	lowercaseMin = 97
	lowercaseMax = 122

	asciiDiff = 32
)

func main() {
	data, err := ioutil.ReadFile("./09_input.txt")
	if err != nil {
		fmt.Println("couldn't read file")
		return
	}

	inputStr := strings.Trim(string(data), "\n")
	newStr := inputStr

	for {
		foundPair := false
		newerStr := ""
		i := 0

		for i < len(newStr) {
			if i == len(newStr)-1 {
				newerStr += string(newStr[i])
				break
			}

			char := newStr[i]
			nextChar := newStr[i+1]

			if lowercase(char) {
				if nextChar == (char - asciiDiff) {
					i += 2
					foundPair = true
				} else {
					newerStr += string(char)
					i++
				}
			} else {
				if nextChar == (char + asciiDiff) {
					i += 2
					foundPair = true
				} else {
					newerStr += string(char)
					i++
				}
			}
		}

		newStr = newerStr

		if !foundPair {
			break
		}
	}

	fmt.Println(len(strings.Trim(newStr, " ")))
}

func lowercase(char byte) bool {
	return char >= lowercaseMin && char <= lowercaseMax
}
