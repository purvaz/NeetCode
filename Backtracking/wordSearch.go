package Backtracking

func Exist(board [][]byte, word string) bool {

	rows := len(board)
	columns := len(board[0])

	var dfs func(r, c, current int) bool
	dfs = func(r, c, current int) bool {

		if r < 0 || c < 0 || r >= rows || c >= columns || current == len(word) {
			return false
		}

		if board[r][c] != word[current] || board[r][c] == '*' {
			return false
		}

		if current == len(word)-1 {
			return true
		}

		temp := board[r][c]
		board[r][c] = '*'

		result := dfs(r+1, c, current+1) || dfs(r-1, c, current+1) || dfs(r, c-1, current+1) || dfs(r, c+1, current+1)

		board[r][c] = temp

		return result
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
