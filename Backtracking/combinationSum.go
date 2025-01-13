package Backtracking

func CombinationSum(candidates []int, target int) [][]int {

	result := make([][]int, 0)
	combination := make([]int, 0)

	var backtrack func(index, currentSum int, combination []int)
	backtrack = func(index, currentSum int, combination []int) {

		if currentSum == target {
			result = append(result, append([]int{}, combination...))
			return
		}
		if currentSum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			combination = append(combination, candidates[i])
			backtrack(i, currentSum+candidates[i], combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(0, 0, combination)
	return result
}
