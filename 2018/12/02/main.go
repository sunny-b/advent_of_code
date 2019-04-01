package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

var (
	dist          = os.Getenv("MAX_DISTANCE")
	inputFilePath = os.Getenv("INPUT_FILE")

	numOfSafeSpots int
	maxWidth       int
	maxDistance    int
)

type coord struct {
	x int
	y int
}

func main() {
	fmt.Println("reading input string")

	absPath, _ := filepath.Abs(inputFilePath)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("couldn't read file")
		return
	}

	var (
		inputStr    = strings.Split(strings.Trim(string(data), "\n"), "\n")
		coordinates = []coord{}
	)

	maxDistance, err := strconv.Atoi(dist)
	if err != nil {
		fmt.Println("could not parse string")
	}

	fmt.Println("parsing coordinates")

	for _, coordStr := range inputStr {
		coords := strings.Split(coordStr, ", ")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			fmt.Println("failed to convert x string")
			return
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			fmt.Println("failed to convert y string")
			return
		}

		if x > maxWidth {
			maxWidth = x
		}
		if y > maxWidth {
			maxWidth = y
		}

		coordinates = append(coordinates, coord{
			x: y,
			y: x,
		})
	}

	fmt.Println("filling grid")

	for i := 0; i < maxWidth; i++ {
		for j := 0; j < maxWidth; j++ {
			if safeDistance(i, j, maxDistance, coordinates) {
				numOfSafeSpots++
			}
		}
	}

	fmt.Println(numOfSafeSpots)
}

func safeDistance(i, j, maxDistance int, coordinates []coord) bool {
	var totalDist int

	for _, c := range coordinates {
		deltaX := int(math.Abs(float64(c.x - i)))
		deltaY := int(math.Abs(float64(c.y - j)))

		totalDist += (deltaX + deltaY)
	}

	return totalDist < maxDistance
}
