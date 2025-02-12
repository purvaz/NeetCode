package DynamicProgramming2D

import "sort"

func change(amount int, coins []int) int {

	sort.Ints(coins)

	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(index, amount int) int
	dfs = func(index, amount int) int {

		if amount == 0 {
			return 1
		}

		if index >= len(coins) {
			return 0
		}

		if dp[index][amount] != -1 {
			return dp[index][amount]
		}

		result := 0
		if amount >= coins[index] {
			result = dfs(index+1, amount)
			result += dfs(index, amount-coins[index])
		}
		dp[index][amount] = result
		return result
	}
	return dfs(0, amount)
}
