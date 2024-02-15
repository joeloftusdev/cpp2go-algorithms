package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Data structure to store a graph edge
type Edge struct {
	source, dest, weight int
}

// Data structure to store a heap node
type Node struct {
	vertex, weight int
}

// A struct to represent a graph object
type Graph struct {
	adjList [][]Edge
}

// Graph Constructor func
func NewGraph(edges []Edge, n int) *Graph {
	graph := &Graph{}

	graph.adjList = make([][]Edge, n)

	for _, edge := range edges {
		graph.adjList[edge.source] = append(graph.adjList[edge.source], edge)
	}
	return graph
}

func printPath(prev []int, i, source int) {
	if i < 0 {
		return
	}
	printPath(prev, prev[i], source)
	if i != source {
		fmt.Printf(", ")
	}
	fmt.Printf("%d", i)

}

// Comparison object to be used to order the heap
type comp []Node

func (c comp) Len() int           { return len(c) }
func (c comp) Less(i, j int) bool { return c[i].weight < c[j].weight }
func (c comp) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c *comp) Push(x interface{}) {
	*c = append(*c, x.(Node))
}
func (c *comp) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

func findShortestPaths(g *Graph, source, n int) {
	// create a min-heap and push source node having distance 0
	minHeap := &comp{}
	heap.Init(minHeap)
	heap.Push(minHeap, Node{source, 0})

	// set initial distance from the source to `v` as infinity
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}

	// distance from the source to itself is zero
	dist[source] = 0

	// boolean array to track vertices for which minimum
	// cost is already found
	done := make([]bool, n)
	done[source] = true

	// stores predecessor of a vertex (to a print path)
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	// run till min-heap is empty

	for len(*minHeap) > 0 {
		// Remove and return the best vertex
		node := heap.Pop(minHeap).(Node)

		// get the vertex number
		u := node.vertex

		// do for each neighbor `v` of `u`
		for _, edge := range g.adjList[u] {
			v := edge.dest
			weight := edge.weight

			// Relaxation step
			if !done[v] && (dist[u]+weight) < dist[v] {
				dist[v] = dist[u] + weight
				prev[v] = u
				heap.Push(minHeap, Node{v, dist[v]})
			}
		}
		// mark vertex `u` as done so it will not get picked up again
		done[u] = true
	}

	for i := 0; i < n; i++ {
		if i != source && dist[i] != math.MaxInt64 {
			fmt.Printf("Path (%d --> %d): Minimum cost = %d, Route = [", source, i, dist[i])
			printPath(prev, i, source)
			fmt.Println("]")
		}
	}
}

func main() {
	edges := []Edge{
		{0, 1, 10}, {0, 4, 3}, {1, 2, 2}, {1, 4, 4}, {2, 3, 9},
		{3, 2, 7}, {4, 1, 1}, {4, 2, 8}, {4, 3, 2},
	}

	n := 5

	graph := NewGraph(edges, n)

	for source := 0; source < n; source++ {
		findShortestPaths(graph, source, n)
	}
}
