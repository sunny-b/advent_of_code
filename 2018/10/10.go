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

	globalMin := -1
	inputStr := strings.Trim(string(data), "\n")

	for j := lowercaseMin; j <= lowercaseMax; j++ {
		lowerChar := j
		upperChar := lowerChar - asciiDiff

		newStr := strings.Replace(inputStr, string(lowerChar), "", -1)
		newStr = strings.Replace(newStr, string(upperChar), "", -1)

		fmt.Println(string(lowerChar), string(upperChar))

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

		stringLen := len(strings.Trim(newStr, " "))

		if globalMin < 0 || stringLen < globalMin {
			globalMin = stringLen
		}
	}

	fmt.Println(globalMin)
}

func lowercase(char byte) bool {
	return char >= lowercaseMin && char <= lowercaseMax
}
