package Backtracking

func FindWord(board [][]byte, word string) bool {

	rows := len(board)
	cols := len(board[0])

	var backtrack func(r, c, currentWord int) bool
	backtrack = func(r, c, currentWord int) bool {

		if r < 0 || c < 0 || r > rows || c > cols || currentWord == len(word) {
			return false
		}

		if board[r][c] != word[currentWord] || board[r][c] == '*' {
			return false
		}

		if board[r][c] == word[currentWord] || currentWord == len(word)-1 {
			return true
		}

		temp := board[r][c]
		board[r][c] = '*'

		result := backtrack(r-1, c, currentWord+1) || backtrack(r+1, c, currentWord+1) || backtrack(r, c-1, currentWord+1) || backtrack(r, c+1, currentWord+1)

		board[r][c] = temp

		return result
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}

	return false

}
