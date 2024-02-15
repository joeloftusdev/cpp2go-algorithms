package main

import (
	"fmt"
)

// A struct to represent a disjoint set
type DisjointSet struct {
	parent map[int]int

	rank map[int]int
}

func (ds *DisjointSet) MakeSet(universe []int) {

	ds.parent = make(map[int]int)
	ds.rank = make(map[int]int)

	for _, item := range universe {
		ds.parent[item] = item
		ds.rank[item] = 0
	}

}

// Find the root of the set in which element `k` belongs
func (ds *DisjointSet) Find(k int) int {
	// if `k` is not the root
	if ds.parent[k] != k {
		// path compression
		ds.parent[k] = ds.Find(ds.parent[k])
	}
	return ds.parent[k]
}

// Perform Union of two subsets
func (ds *DisjointSet) Union(a, b int) {
	// find the root of the sets in which elements `x` and `y` belongs
	x := ds.Find(a)
	y := ds.Find(b)

	// if `x` and `y` are present in the same set
	if x == y {
		return
	}

	// Always attach a smaller depth tree under the root of the deeper tree.
	if ds.rank[x] > ds.rank[y] {
		ds.parent[y] = x
	} else if ds.rank[x] < ds.rank[y] {
		ds.parent[x] = y
	} else {
		ds.parent[x] = y
		ds.rank[y]++
	}
}

func printSets(universe []int, ds DisjointSet) {
	for _, item := range universe {
		fmt.Printf("%d ", ds.Find(item))
	}
	fmt.Println()
}

// Disjoint–Set data structure (Union–Find algorithm)
func main() {

	//universe items
	universe := []int{1, 2, 3, 4, 5}

	//init DisjointSet
	ds := DisjointSet{}

	// create a singleton set for each element of the universe
	ds.MakeSet(universe)
	printSets(universe, ds)

	ds.Union(4, 3) // 4 and 3 are in the same set
	printSets(universe, ds)

	ds.Union(2, 1) // 1 and 2 are in the same set
	printSets(universe, ds)

	ds.Union(1, 3) // 1, 2, 3, 4 are in the same set
	printSets(universe, ds)
}
