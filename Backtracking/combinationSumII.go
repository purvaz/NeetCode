package Backtracking

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {

	result := make([][]int, 0)
	combination := make([]int, 0)

	sort.Ints(candidates)

	var backtrack func(idx, sum int, combination []int)
	backtrack = func(idx, sum int, combination []int) {

		if sum == target {
			result = append(result, append([]int{}, combination...))
			return
		}

		if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i] == candidates[i-1] {
				continue
			}

			combination = append(combination, candidates[i])
			backtrack(i, sum+candidates[i], combination)
			combination = combination[:len(candidates)-1]
		}
	}
	backtrack(0, 0, combination)
	return result
}
