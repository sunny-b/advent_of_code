package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	stringRegex = regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")

	nodeMap   = make(map[string]*treeNode)
	taskOrder []string
)

type treeNode struct {
	inDegree  int
	completed bool
	children  []string
}

func main() {
	absPath, _ := filepath.Abs("../13_input.txt")
	fmt.Println(findTaskOrder(absPath))
}

func findTaskOrder(filepath string) string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	strSlice := strings.Split(strings.Trim(string(data), "\n"), "\n")

	for _, s := range strSlice {
		matches := stringRegex.FindStringSubmatch(s)

		node, ok := nodeMap[matches[1]]

		var newNode *treeNode
		if !ok {
			newNode = &treeNode{
				inDegree: 0,
				children: append([]string{}, matches[2]),
			}
		} else {
			newNode = node
			newNode.children = append(node.children, matches[2])
		}

		nodeMap[matches[1]] = newNode

		node, ok = nodeMap[matches[2]]

		if !ok {
			newNode = &treeNode{
				inDegree: 1,
			}
		} else {
			newNode = node
			newNode.inDegree++
		}

		nodeMap[matches[2]] = newNode
	}

	for {
		var (
			found  bool
			minVal = "z"
		)

		printNodes(nodeMap)

		for key, node := range nodeMap {
			if !node.completed && node.inDegree == 0 && key < minVal {
				minVal = key
				found = true
			}
		}

		if !found {
			break
		}

		taskOrder = append(taskOrder, minVal)

		nodeMap[minVal].completed = true
		for _, child := range nodeMap[minVal].children {
			nodeMap[child].inDegree--
		}
	}

	return strings.Join(taskOrder, "")
}

func printNodes(nodeMap map[string]*treeNode) {
	for k, v := range nodeMap {
		fmt.Println(k, *v)
	}
}
