package Heaps

import (
	"container/heap"
)

func FindKthLargest(nums []int, k int) int {

	for i := 0; i < len(nums); i++ {
		nums[i] *= -1
	}
	values := IntHeap(nums)
	heap.Init(&values)

	for i := 0; i < k-1; i++ {
		heap.Pop(&values)
	}
	return values[0] * -1
}
