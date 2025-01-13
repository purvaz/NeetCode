package DynamicProgramming

func MaxProduct(nums []int) int {

	result := nums[0]
	currMin, currMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max(nums[i]*currMax, nums[i])
		currMin = min(nums[i]*currMin, nums[i])
		result = max(currMax, result)
	}
	return result
}
