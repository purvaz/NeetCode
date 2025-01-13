package Backtracking

import (
	"sort"
)

func SubsetsWithDup(nums []int) [][]int {

	result := make([][]int, 0)
	subset := make([]int, 0)

	sort.Ints(nums)

	var backtrack func(index int)
	backtrack = func(index int) {

		result = append(result, append([]int{}, subset...))
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			subset = append(subset, nums[i])
			backtrack(i + 1)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack(0)
	return result
}
