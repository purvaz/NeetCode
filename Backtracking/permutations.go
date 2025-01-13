package Backtracking

func Permute(nums []int) [][]int {

	result := make([][]int, 0)
	permutation := make([]int, 0)
	// keep a track of the elements and count of its occurrence
	count := make(map[int]int)

	var backtrack func(index int)
	backtrack = func(index int) {

		if len(permutation) == len(nums) {
			result = append(result, append([]int{}, permutation...))
		}

		for i := 0; i < len(nums); i++ {
			// this means it has not been considered yet
			if count[i] == 0 {
				count[i]++
				permutation = append(permutation, nums[i])
				backtrack(i + 1)
				permutation = permutation[:len(permutation)-1]
				count[i]--
			}
		}
	}
	backtrack(0)
	return result
}
