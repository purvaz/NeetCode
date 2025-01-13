package DynamicProgramming

func LengthOfLIS(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	lis := make([]int, len(nums))
	maxLIS := 0

	lis[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		maxLen := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && maxLen < (1+lis[j]) {
				maxLen = 1 + lis[j]
			}
		}
		lis[i] = max(1, maxLen)
		if lis[i] > maxLIS {
			maxLIS = lis[i]
		}
	}

	return maxLIS
}
