package Heaps

import "container/heap"

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func KConstructor(k int, nums []int) KthLargest {

	// creating a list of type IntHeap
	values := IntHeap(nums)

	// initialize the object of type struct, passing it the list of type IntHeap and k
	this := KthLargest{minHeap: &values, k: k}

	// initialize the min-heap. this will have all the elements in the beginning
	heap.Init(this.minHeap)

	// now, we pop all the elements till len(heap) becomes = k
	for len(*this.minHeap) > k {
		heap.Pop(this.minHeap)
	}
	return this
}

func (this *KthLargest) Add(val int) int {

	// push the value in the heap
	heap.Push(this.minHeap, val)

	// if the length of the heap is > k, then pop the smallest element
	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}
