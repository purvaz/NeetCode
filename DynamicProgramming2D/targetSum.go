package DynamicProgramming2D

func findTargetSumWays(nums []int, target int) int {

	dp := make([]map[int]int, len(nums)+1)

	for i := 0; i <= len(nums); i++ {
		dp[i] = make(map[int]int)
	}

	dp[0][0] = 1

	for i := 0; i < len(nums); i++ {
		for sum, count := range dp[i] {
			dp[i+1][sum+nums[i]] += count
			dp[i+1][sum-nums[i]] += count
		}
	}
	return dp[len(nums)][target]
}
