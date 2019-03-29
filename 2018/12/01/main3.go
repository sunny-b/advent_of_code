package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

var (
	infiniteMarkers = make(map[string]bool)
	markerAreas     = make(map[string]int)
	markers         = make(map[string]bool)

	maxWidth    int
	maxDistance int
)

type coord struct {
	x    int
	y    int
	mark string
}

func main() {
	fmt.Println("reading input string")

	absPath, _ := filepath.Abs("../12_input.txt")
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("couldn't read file")
		return
	}

	var (
		inputStr    = strings.Split(strings.Trim(string(data), "\n"), "\n")
		coordinates = []coord{}
	)

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

	fmt.Println("placing markers in matrix")

	for i, c := range coordinates {
		char := string(alphabet[(i % len(alphabet))])
		id := char + strconv.Itoa(i)

		c.mark = strings.ToUpper(id)

		markers[c.mark] = true

		coordinates[i] = c
	}

	fmt.Println("filling grid")

	for i := 0; i < maxWidth; i++ {
		for j := 0; j < maxWidth; j++ {
			marker := strings.ToLower(nearestMarker(i, j, coordinates))

			if i == 0 || j == 0 || i == maxWidth-1 || j == maxWidth-1 {
				infiniteMarkers[marker] = true
			}

			if !infiniteMarkers[marker] {
				markerAreas[marker]++
			}
		}
	}

	fmt.Println("finding maximum area")

	var maxArea int
	for _, area := range markerAreas {
		if area > maxArea {
			maxArea = area
		}
	}

	fmt.Println(maxArea)
}

func nearestMarker(i, j int, coordinates []coord) string {
	var (
		currentMin int
		globalMins []coord
	)

	for _, c := range coordinates {
		deltaX := int(math.Abs(float64(c.x - i)))
		deltaY := int(math.Abs(float64(c.y - j)))

		if len(globalMins) == 0 || (deltaX+deltaY) == currentMin {
			currentMin = (deltaX + deltaY)
			globalMins = append(globalMins, c)
		} else if (deltaX + deltaY) < currentMin {
			currentMin = (deltaX + deltaY)
			globalMins = append([]coord{}, c)
		}
	}

	if len(globalMins) > 1 {
		return "."
	}

	return globalMins[0].mark
}
