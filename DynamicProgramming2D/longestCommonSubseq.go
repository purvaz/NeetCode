package DynamicProgramming2D

func LongestCommonSubsequence(text1 string, text2 string) int {

	grid := make([][]int, len(text1)+1)
	for i := range grid {
		grid[i] = make([]int, len(text2)+1)
	}

	for i := len(grid) - 2; i >= 0; i-- {
		for j := len(grid[0]) - 2; j >= 0; j-- {
			if text1[i] == text2[j] {
				grid[i][j] = 1 + grid[i+1][j+1]
			} else {
				grid[i][j] = max(grid[i+1][j], grid[i][j+1])
			}
		}
	}

	return grid[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
