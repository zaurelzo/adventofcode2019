package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

//build a bidirectional graph if asked
func buildGraph(graphFileName string, bidirectional bool) (map[string][]string, string, string) {
	file, err := os.Open(graphFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	elts := strings.Split(string(content), "\n")
	graph := make(map[string][]string)
	var youOrbital string
	var sanOrbital string
	for _, graphElement := range elts {
		nodes := strings.Split(graphElement, ")")
		if len(nodes) != 2 {
			panic("invalid number of nodes")
		}
		graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
		if bidirectional {
			graph[nodes[1]] = append(graph[nodes[1]], nodes[0])
		}
		if nodes[1] == "YOU" {
			youOrbital = nodes[0]
		}

		if nodes[1] == "SAN" {
			sanOrbital = nodes[0]
		}

	}
	return graph, youOrbital, sanOrbital
}

func totalNbOfOrbits(graph map[string][]string) int {
	numberOforbits := 0
	for node := range graph {
		numberOforbits += nbOrbitsFromNode(graph, node)
	}
	return numberOforbits
}

// do a DFS to count the total edges from initialNode
func nbOrbitsFromNode(graph map[string][]string, initialNode string) int {
	numberOforbit := 0
	visitedNodes := make(map[string]bool)
	stack := make([]string, 0)
	stack = append(stack, initialNode)
	for len(stack) != 0 {
		//pop :
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited, exist := visitedNodes[currentNode]
		if !exist || !visited {
			visitedNodes[currentNode] = true
			for _, nextNode := range graph[currentNode] {
				numberOforbit++
				stack = append(stack, nextNode)
			}
		}

	}
	return numberOforbit
}

func printGraph(graph map[string][]string) {
	for nodeName, adjacentNode := range graph {
		fmt.Printf(" %v -> %v \n", nodeName, adjacentNode)
	}
}

// apply  Dijkstra algorithm to find shortest path between source and dest
func minimalNbOfOrbitalToTransfer(graph map[string][]string, source string, dest string) int {
	set := make(map[string]bool) //simulate the set data struct since golang does not have the built-in data structure
	dist := make(map[string]int)
	set[source] = true
	// Initializations
	dist[source] = 0
	for nodeName := range graph {
		if nodeName != source {
			dist[nodeName] = math.MaxInt32 //INFINITY VALUE
			set[nodeName] = true
		}
	}

	//algo
	for len(set) != 0 {
		nodeWithSmallestDist := findNodeWithMinDist(set, dist)
		delete(set, nodeWithSmallestDist)
		for _, neighbor := range graph[nodeWithSmallestDist] {
			_, neighborInSet := set[neighbor]
			if neighborInSet {
				currentDist := dist[nodeWithSmallestDist] + 1
				if currentDist < dist[neighbor] {
					dist[neighbor] = currentDist
				}
			}
		}
	}
	return dist[dest]
}

func findNodeWithMinDist(set map[string]bool, distances map[string]int) string {
	var min *int
	var nodeWithSmallestDist string
	for node := range set {
		nodeDist, exist := distances[node]
		if exist {
			if min == nil || (nodeDist < *min) {
				min = &nodeDist
				nodeWithSmallestDist = node
			}
		}

	}
	return nodeWithSmallestDist
}

func main() {
	graph, _, _ := buildGraph("input-graph", false)
	//graph := buildGraph("graph", false)
	//graph := buildGraph("simpleGraph",false)
	//printGraph(graph)
	fmt.Printf("Part one, nbOfOrbit  %d \n", totalNbOfOrbits(graph))
	biDirectionalGraph, you, san := buildGraph("input-graph", true)
	fmt.Printf("Part two, minimum number of orbital transfers %+v", minimalNbOfOrbitalToTransfer(biDirectionalGraph, you, san))
}
