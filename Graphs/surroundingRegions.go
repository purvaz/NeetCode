package Graphs

func solve(board [][]byte) {

	rows, cols := len(board), len(board[0])

	var capture func(r, c int)
	capture = func(r, c int) {

		if r < 0 || c < 0 || r == rows || c == cols || board[r][c] != 0 {
			return
		}

		board[r][c] = 'T'
		capture(r-1, c)
		capture(r+1, c)
		capture(r, c-1)
		capture(r, c+1)
	}

	for r := 0; r < rows; r++ {
		if board[r][0] == 'O' {
			capture(r, 0)
		}
		if board[r][cols-1] == 'O' {
			capture(r, cols-1)
		}
	}

	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' {
			capture(0, c)
		}
		if board[rows-1][c] == 'O' {
			capture(rows-1, c)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
			if board[r][c] == 'T' {
				board[r][c] = 'O'
			}
		}
	}

}
