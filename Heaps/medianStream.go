package Heaps

import (
	"container/heap"
)

type MedianFinder struct {
	small IntHeap
	big   IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{small: IntHeap{}, big: IntHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	// push in big or small depending on num > mid or num < mid
	if len(this.big) > 0 && num > this.big[0] {
		heap.Push(&this.big, num)
	} else {
		heap.Push(&this.small, (num * -1))
	}

	// swap heaps if len of one is > than other by 2
	// basically this makes sure that the diff between the len of two heaps is <= 1
	if len(this.big) > len(this.small)+1 {
		number := heap.Pop(&this.big).(int)
		heap.Push(&this.small, (number * -1))
	} else if len(this.small) > len(this.big)+1 {
		number := heap.Pop(&this.small).(int)
		heap.Push(&this.big, (number * -1))
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if len(this.big) > len(this.small) {
		return heap.Pop(&this.big).(float64)
	} else if len(this.small) > len(this.big) {
		return heap.Pop(&this.small).(float64)
	}
	return float64(this.big[0]+(this.small[0]*-1)) / 2
}
