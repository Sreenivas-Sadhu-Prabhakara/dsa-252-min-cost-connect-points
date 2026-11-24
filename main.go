package main

import "fmt"

// MinCostConnect returns the cheapest total wire connecting all points,
// edges priced by Manhattan distance — an MST on an implicit complete
// graph. Dense Prim without a heap: track each outside point's cheapest
// link to the tree and absorb the cheapest per round.
func MinCostConnect(points [][2]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}
	const inf = int(^uint(0) >> 1)
	dist := make([]int, n) // cheapest edge from the tree to each point
	inTree := make([]bool, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0
	total := 0
	for round := 0; round < n; round++ {
		best := -1
		for v := 0; v < n; v++ {
			if !inTree[v] && (best == -1 || dist[v] < dist[best]) {
				best = v
			}
		}
		inTree[best] = true
		total += dist[best]
		for v := 0; v < n; v++ {
			if !inTree[v] {
				d := abs(points[best][0]-points[v][0]) + abs(points[best][1]-points[v][1])
				if d < dist[v] {
					dist[v] = d
				}
			}
		}
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(MinCostConnect([][2]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
}
