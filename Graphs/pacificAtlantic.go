package Graphs

func pacificAtlantic(heights [][]int) [][]int {

	rows, cols := len(heights), len(heights[0])
	pacific := make(map[[2]int]bool)
	atlantic := make(map[[2]int]bool)

	var dfs func(r, c int, visited map[[2]int]bool, prevHeight int)
	dfs = func(r, c int, visited map[[2]int]bool, prevHeight int) {

		element := [2]int{r, c}
		if visited[element] ||
			r < 0 || c < 0 ||
			r >= rows || c >= cols ||
			heights[r][c] < prevHeight {
			return
		}

		visited[element] = true

		dfs(r-1, c, visited, heights[r][c])
		dfs(r+1, c, visited, heights[r][c])
		dfs(r, c-1, visited, heights[r][c])
		dfs(r, c+1, visited, heights[r][c])
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, pacific, heights[0][c])
		dfs(rows-1, c, pacific, heights[rows-1][c])
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, atlantic, heights[r][0])
		dfs(r, cols-1, atlantic, heights[r][cols-1])
	}

	result := make([][]int, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			element := [2]int{r, c}
			if pacific[element] && atlantic[element] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
