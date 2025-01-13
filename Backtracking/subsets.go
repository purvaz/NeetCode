package Backtracking

func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	subset := make([]int, 0)

	var backtrack func(index int)
	backtrack = func(index int) {
		result = append(result, append([]int{}, subset...))
		if index == len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			subset = append(subset, nums[i])
			backtrack(i + 1)
			subset = subset[:len(subset)-1]

		}
	}
	backtrack(0)
	return result
}
