package Graphs

func MaxAreaOfIsland(grid [][]int) int {

	rows, cols := len(grid), len(grid[0])
	maxArea := 0
	area := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return
		}

		grid[r][c] = 0
		area += 1

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				area = 0
				dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
