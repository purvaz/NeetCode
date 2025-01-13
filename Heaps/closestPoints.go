package Heaps

import (
	"container/heap"
	"math"
)

type Points struct {
	dist int
	x    int
	y    int
}

type PointHeap []*Points

func (h PointHeap) Len() int { return len(h) }

// Since the comparison is based on distance, the Less function is modified
func (h PointHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }

func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// We want to push the entire object, and not a value, so typecasting to (*Points)
func (h *PointHeap) Push(x any) {
	*h = append(*h, x.(*Points))
}

func (h *PointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KClosest(points [][]int, k int) [][]int {

	result := make([][]int, 0)
	_points := make([]*Points, 0, len(points))
	for _, point := range points {
		// avoiding math.Sqrt on the whole equation
		// as taking the sqrt of the values would not change the relative order
		distance := math.Sqrt(math.Pow(math.Abs(float64(point[0]-0)), 2) + math.Pow(math.Abs(float64(point[1]-0)), 2))
		_points = append(_points, &Points{dist: int(distance), x: point[0], y: point[1]})
	}

	// create a heap of type PointHeap, and initialize it
	pts := PointHeap(_points)
	heap.Init(&pts)

	// pop till we have k values
	for i := 0; i < k; i++ {
		temp := heap.Pop(&pts).(*Points)
		result = append(result, []int{temp.x, temp.y})
	}
	return result
}
