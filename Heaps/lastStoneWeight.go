package Heaps

import (
	"container/heap"
	"math"
)

func LastStoneWeight(stones []int) int {

	// convert all values by multiplying them by -1
	for i := 0; i < len(stones); i++ {
		stones[i] *= -1
	}

	// initialize the heap
	_stones := IntHeap(stones)
	heap.Init(&_stones)

	// check the conditions by popping two largest stones,
	// and push the difference if stone1 != stone2
	for len(_stones) > 1 {
		stone1, _ := heap.Pop(&_stones).(int)
		stone2, _ := heap.Pop(&_stones).(int)
		if stone1 != stone2 {
			heap.Push(&_stones, (stone1 - stone2))
		}
	}
	// in the end, the array might be empty.
	// in such case, we just append a 0 and return 1st value
	_stones = append(_stones, 0)
	return int(math.Abs(float64(_stones[0])))
}
