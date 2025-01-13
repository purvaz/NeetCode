package DynamicProgramming2D

func uniquePaths(m int, n int) int {

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			if row == m-1 && col == n-1 {
				grid[row][col] = 1
			} else if row == m-1 {
				grid[row][col] = grid[row][col+1]
			} else if col == n-1 {
				grid[row][col] = grid[row+1][col]
			} else {
				grid[row][col] = grid[row][col+1] + grid[row+1][col]
			}
		}
	}
	return grid[0][0]
}
