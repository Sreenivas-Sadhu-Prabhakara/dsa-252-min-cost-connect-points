package main

import "testing"

func TestMinCostConnect(t *testing.T) {
	cases := []struct {
		points [][2]int
		want   int
	}{
		{[][2]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
		{[][2]int{{3, 12}, {-2, 5}, {-4, 1}}, 18},
		{[][2]int{{0, 0}}, 0},
		{nil, 0},
		{[][2]int{{0, 0}, {1, 1}}, 2},
		{[][2]int{{0, 0}, {0, 1}, {0, 2}}, 2}, // chain along an axis
	}
	for _, c := range cases {
		if got := MinCostConnect(c.points); got != c.want {
			t.Errorf("MinCostConnect(%v) = %d, want %d", c.points, got, c.want)
		}
	}
}
