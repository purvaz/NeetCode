package Intervals

func insert(intervals [][]int, newInterval []int) [][]int {

	result := [][]int{}

	for index, interval := range intervals {
		if newInterval[1] < interval[0] {
			result = append(result, newInterval)
			return append(result, intervals[index:]...)
		} else if newInterval[0] > interval[1] {
			result = append(result, interval)
		} else {
			newInterval = []int{min(newInterval[0], interval[0]), max(newInterval[1], interval[1])}
		}
	}
	result = append(result, newInterval)
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
