package main

import (
	"fmt"
	"sort"
)

// Data structure to store a graph edge
type Edge struct {
	src, dest, weight int
}

// Comparison object to be used to order the edges
type ByWeight []Edge

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].weight > a[j].weight }

// A struct to represent a disjoint set
type DisjointSet struct {
	parent map[int]int
}

// perform MakeSet operation
func (ds *DisjointSet) MakeSet(n int) {
	ds.parent = make(map[int]int)
	// create `n` disjoint sets (one for each vertex)
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}
}

// Find the root of the set in which element `k` belongs
func (ds *DisjointSet) Find(k int) int {
	if ds.parent[k] == k {
		return k
	}

	// recur for the parent until we find the root
	return ds.Find(ds.parent[k])
}

// Perform Union of two subsets
func (ds *DisjointSet) Union(a, b int) {
	// find the root of the sets in which elements
	// `x` and `y` belongs
	x := ds.Find(a)
	y := ds.Find(b)

	ds.parent[x] = y

}

// Function to construct MST using Kruskal’s algorithm
func runKruskalAlgorithm(edges []Edge, n int) []Edge {

	// stores the edges present in MST
	var MST []Edge

	// initialize `DisjointSet` class
	var ds DisjointSet

	// create a singleton set for each element of the universe
	ds.MakeSet(n)

	// sort edges by increasing weight
	sort.Sort(ByWeight(edges))

	// MST contains exactly `V-1` edges
	for len(MST) != n-1 {
		// consider the next edge with minimum weight from the graph
		nextEdge := edges[len(edges)-1]
		edges = edges[:len(edges)-1]

		// find the root of the sets to which two endpoints
		// vertices of the next edge belongs
		x := ds.Find(nextEdge.src)
		y := ds.Find(nextEdge.dest)

		// if both endpoints have different parents, they belong to
		// different connected components and can be included in MST
		if x != y {
			MST = append(MST, nextEdge)
			ds.Union(x, y)
		}
	}
	return MST
}

func main() {

	// vector of graph edges as per the above diagram.
	edges := []Edge{
		// (u, v, w) triplet represent undirected edge from
		// vertex `u` to vertex `v` having weight `w`
		{0, 1, 7}, {1, 2, 8}, {0, 3, 5}, {1, 3, 9}, {1, 4, 7}, {2, 4, 5},
		{3, 4, 15}, {3, 5, 6}, {4, 5, 8}, {4, 6, 9}, {5, 6, 11},
	}

	// total number of nodes in the graph (labelled from 0 to 6)
	n := 7

	// construct graph
	edges = runKruskalAlgorithm(edges, n)

	for _, edge := range edges {
		fmt.Printf("(%d, %d, %d)\n", edge.src, edge.dest, edge.weight)
	}
}
