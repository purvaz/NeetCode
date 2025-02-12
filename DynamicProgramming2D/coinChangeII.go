package DynamicProgramming2D

import "sort"

func change(amount int, coins []int) int {

	sort.Ints(coins)

	// Initialize the 2D array for memoization
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	// Design a DFS function that accepts the current index and the remaining amount
	// as parameters and returns an int as the number of ways in which
	// this amount can be constructed.
	var dfs func(index, amount int) int
	dfs = func(index, amount int) int {

		// if the amount is 0, it means it is one of the solutions.
		if amount == 0 {
			return 1
		}

		// check if the index exceeds the len(coins), if so, return 0.
		if index >= len(coins) {
			return 0
		}

		// if the current value is not -1, return the value at that cell.
		if dp[index][amount] != -1 {
			return dp[index][amount]
		}

		result := 0
		// If the amount is >= the coin face value
		if amount >= coins[index] {
			// recursively call the dfs function for neighboring value
			result = dfs(index+1, amount)
			// compute the result by adding the value returned by the dfs function
			result += dfs(index, amount-coins[index])
		}
		// Assign the updated result value in the 2D array
		dp[index][amount] = result
		return result
	}
	return dfs(0, amount)
}
