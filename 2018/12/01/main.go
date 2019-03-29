package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// const (
// 	alphabet = "abcdefghijklmnopqrstuvwxyz"
// )

// var (
// 	infiniteMarkers = make(map[string]bool)
// 	markerAreas     = make(map[string]int)
// 	markers         = make(map[string]bool)

// 	maxWidth    int
// 	maxDistance int
// )

// type graph [][]string

// type coord struct {
// 	x int
// 	y int
// }

// func main() {
// 	start := time.Now()
// 	fmt.Println("reading input string")

// 	absPath, _ := filepath.Abs("../12_input.txt")
// 	data, err := ioutil.ReadFile(absPath)
// 	if err != nil {
// 		fmt.Println("couldn't read file")
// 		return
// 	}

// 	var (
// 		inputStr    = strings.Split(strings.Trim(string(data), "\n"), "\n")
// 		coordinates = []coord{}
// 	)

// 	fmt.Println("parsing coordinates")

// 	for _, coordStr := range inputStr {
// 		coords := strings.Split(coordStr, ", ")

// 		x, err := strconv.Atoi(coords[0])
// 		if err != nil {
// 			fmt.Println("failed to convert x string")
// 			return
// 		}

// 		y, err := strconv.Atoi(coords[1])
// 		if err != nil {
// 			fmt.Println("failed to convert y string")
// 			return
// 		}

// 		if x > maxWidth {
// 			maxWidth = x
// 		}
// 		if y > maxWidth {
// 			maxWidth = y
// 		}

// 		coordinates = append(coordinates, coord{
// 			x: y,
// 			y: x,
// 		})
// 	}

// 	fmt.Println("calculating max distance")

// 	for i, c := range coordinates {
// 		var currentMax int

// 		for j := i + 1; j < len(coordinates); j++ {
// 			next := coordinates[j]
// 			deltaX := int(math.Abs(float64(c.x - next.x)))
// 			deltaY := int(math.Abs(float64(c.y - next.y)))

// 			if (deltaX + deltaY) > currentMax {
// 				currentMax = (deltaX + deltaY)
// 			}
// 		}

// 		if currentMax > maxDistance {
// 			maxDistance = currentMax
// 		}
// 	}

// 	maxDistance = 7

// 	fmt.Println("max distance: ", maxDistance)

// 	fmt.Println("creating matrix")

// 	g := make(graph, maxWidth+1)
// 	for i := range g {
// 		g[i] = make([]string, maxWidth+1)
// 	}

// 	fmt.Println("placing markers in matrix")

// 	for i, c := range coordinates {
// 		char := string(alphabet[(i % len(alphabet))])
// 		id := char + strconv.Itoa(i)

// 		g[c.x][c.y] = strings.ToUpper(id)
// 		markers[strings.ToUpper(id)] = true
// 	}

// 	fmt.Println("calculating areas of each marker")

// for i, row := range g {
// 	for j, mark := range row {
// 		if markers[mark] {
// 			m := strings.ToLower(mark)

// 			if !infiniteMarkers[m] {
// 				markerAreas[m]++
// 			}

// 			continue
// 		}

// 		m, d := nearestMarker(i, j, 0, g, make(map[string]bool))
// 		marker := strings.ToLower(m)

// 		g[i][j] = fmt.Sprintf("%v-%v", marker, d)

// 		if i == 0 || j == 0 || i == len(g)-1 || j == len(g[0])-1 {
// 			infiniteMarkers[marker] = true
// 		}

// 		if !infiniteMarkers[marker] {
// 			markerAreas[marker]++
// 		}
// 	}
// }

// 	fmt.Println("finding maximum area")

// 	var maxArea int
// 	for _, area := range markerAreas {
// 		if area > maxArea {
// 			maxArea = area
// 		}
// 	}

// 	end := time.Since(start)

// 	ttf := end.Seconds()

// 	fmt.Println()
// 	fmt.Println("seconds to finish: ", ttf)
// 	fmt.Println()

// 	fmt.Println(maxArea)
// }

// func nearestMarker(i, j, d int, g graph, visited map[string]bool) (string, int) {
// 	curr := g[i][j]

// 	if markers[curr] {
// 		return curr, d
// 	} else if curr != "" {
// 		splitMark := strings.Split(curr, "-")
// 		dist, _ := strconv.Atoi(splitMark[1])

// 		return strings.ToUpper(splitMark[0]), (dist + d)
// 	} else if d > maxDistance {
// 		return "", d
// 	}

// 	var (
// 		top, right, left, bottom string
// 		td, rd, ld, bd           int

// 		marker   string
// 		shortest int
// 	)

// 	d++
// 	visited[fmt.Sprintf("%v-%v", i, j)] = true

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("LOGGING START")
// 	// 	fmt.Println("visited: ", visited)
// 	// 	fmt.Println("distance: ", d)
// 	// }

// 	if (j+1) <= maxWidth && !visited[fmt.Sprintf("%v-%v", i, j+1)] {
// 		// if i == 8 && j == 8 {
// 		// 	fmt.Println("yes")
// 		// }
// 		temp := newVisited(visited)
// 		right, rd = nearestMarker(i, j+1, d, g, temp)
// 	}

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("right: ", right)
// 	// 	fmt.Println("rd: ", rd)
// 	// }

// 	if right != "" {
// 		marker = right
// 		shortest = rd
// 	}

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("FIRST CONDITIONAL")
// 	// 	fmt.Println("marker: ", marker)
// 	// 	fmt.Println("shortest: ", shortest)
// 	// }

// 	if (i+1) <= maxWidth && !visited[fmt.Sprintf("%v-%v", i+1, j)] {
// 		// if i == 8 && j == 8 {
// 		// 	fmt.Println("yes")
// 		// }
// 		temp := newVisited(visited)
// 		bottom, bd = nearestMarker(i+1, j, d, g, temp)
// 	}

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("bottom: ", bottom)
// 	// 	fmt.Println("bd: ", bd)
// 	// }

// 	marker, shortest = findShortestMarker(marker, shortest, bottom, bd)

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("SECOND CONDITIONAL")
// 	// 	fmt.Println("marker: ", marker)
// 	// 	fmt.Println("shortest: ", shortest)
// 	// }

// 	if (j-1) >= 0 && !visited[fmt.Sprintf("%v-%v", i, j-1)] {
// 		// if i == 8 && j == 8 {
// 		// 	fmt.Println("yes")
// 		// }
// 		temp := newVisited(visited)
// 		left, ld = nearestMarker(i, j-1, d, g, temp)
// 	}

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("left: ", left)
// 	// 	fmt.Println("ld: ", ld)
// 	// }

// 	marker, shortest = findShortestMarker(marker, shortest, left, ld)

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("THIRD CONDITIONAL")
// 	// 	fmt.Println("marker: ", marker)
// 	// 	fmt.Println("shortest: ", shortest)
// 	// }

// 	if (i-1) >= 0 && !visited[fmt.Sprintf("%v-%v", i-1, j)] {
// 		// if i == 8 && j == 8 {
// 		// 	fmt.Println("yes")
// 		// }
// 		temp := newVisited(visited)
// 		top, td = nearestMarker(i-1, j, d, g, temp)
// 	}

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("top: ", top)
// 	// 	fmt.Println("td: ", td)
// 	// }

// 	marker, shortest = findShortestMarker(marker, shortest, top, td)

// 	// if i == 8 && j == 8 {
// 	// 	fmt.Println("last CONDITIONAL")
// 	// 	fmt.Println("marker: ", marker)
// 	// 	fmt.Println("shortest: ", shortest)
// 	// 	fmt.Println()
// 	// }

// 	// if marker == "" && shortest == 0 {
// 	// 	fmt.Println("FOUND")
// 	// 	fmt.Println("i :", i)
// 	// 	fmt.Println("j :", j)
// 	// 	fmt.Println("d :", d)
// 	// 	fmt.Println("visited: ", visited)
// 	// 	fmt.Println("g: ", g)
// 	// }

// 	return marker, shortest
// }

// func findShortestMarker(oldMark string, oldDist int, newMark string, newDist int) (string, int) {
// 	if newMark == oldMark {
// 		dist := oldDist
// 		if newDist < dist {
// 			dist = newDist
// 		}
// 		return newMark, dist
// 	} else if newDist != 0 {
// 		if oldDist == 0 || newDist < oldDist {
// 			return newMark, newDist
// 		} else if (newDist == oldDist) && (newMark != oldMark) {
// 			return ".", oldDist
// 		}
// 	}

// 	return oldMark, oldDist
// }

// func newVisited(oldVisited map[string]bool) map[string]bool {
// 	newVisit := make(map[string]bool)

// 	for k, v := range oldVisited {
// 		newVisit[k] = v
// 	}

// 	return newVisit
// }

// // [a0-2 a0-1 a0-2 a0-3 a0-4 .-5 c2-5 c2-4 c2-3 c2-4]
// // [a0-1 A0 a0-1 a0-2 a0-3 .-4 c2-4 c2-3 c2-2 c2-3]
// // [a0-2 a0-1 a0-2 d3-2 d3-3 e4-3 c2-3 c2-2 c2-1 c2-2]
// // [a0-3 a0-2 d3-2 d3-1 d3-2 e4-2 c2-2 c2-1 C2 c2-1]
// // [.-3 .-2 d3-1 D3 d3-1 e4-1 e4-2 c2-2 c2-1 c2-2]
// // [b1-2 b1-1 .-2 d3-1 e4-1 E4 e4-1 e4-2 c2-2 c2-3]
// // [b1-1 B1 b1-1 .-2 e4-2 e4-1 e4-2 e4-3 .-3 .-4]
// // [b1-2 b1-1 b1-2 .-3 e4-3 e4-2 e4-3 f5-3 f5-2 f5-3]
// // [b1-3 b1-2 b1-3 .-4 e4-4 e4-3 f5-3 f5-2 f5-1 f5-2]
// // [b1-4 b1-3 b1-4 .-5 f5-4 f5-3 f5-2 f5-1 F5 f5-1]
