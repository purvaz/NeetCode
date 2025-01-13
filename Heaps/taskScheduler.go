package Heaps

import (
	"container/heap"
)

func LeastInterval(tasks []byte, n int) int {

	// convert the byte occurrence to number of times it occurs in the stream
	countMap := make(map[byte]int)
	for _, task := range tasks {
		countMap[task] += 1
	}

	// create an array from the count
	_tasks := make([]int, 0)
	for _, task := range countMap {
		_tasks = append(_tasks, (task * -1))
	}

	// create and initialize heap
	tasksHeap := IntHeap(_tasks)
	heap.Init(&tasksHeap)

	// use a queue to maintain the scheduling of the tasks
	queue := make([][]int, 0)
	time := 0

	// iterate till there are values in either the heap or the queue
	for len(tasksHeap) > 0 || len(queue) > 0 {
		time++
		// if heap is not empty
		if len(tasksHeap) > 0 {
			// pop highest count element from heap
			_count := (heap.Pop(&tasksHeap).(int) + 1)
			// add it to the queue
			if _count != 0 {
				_queue := []int{_count, (time + n)}
				queue = append(queue, _queue)
			}
		}
		// if its time to pop the 1st element from queue
		if len(queue) > 0 && time == queue[0][1] {
			top := queue[0]
			queue = queue[1:]
			// push back to the heap
			heap.Push(&tasksHeap, top[0])
		}
	}
	return time
}
