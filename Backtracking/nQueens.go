package Backtracking

func solveNQueens(n int) [][]string {

	result := make([][]string, 0)
	current := make([]string, 0)

	column := make(map[int]bool)
	positiveDiag := make(map[int]bool)
	negativeDiag := make(map[int]bool)

	var backtrack func(r int)
	backtrack = func(r int) {

		if r == n {
			result = append(result, append([]string{}, current...))
			return
		}

		for c := 0; c < n; c++ {
			// already present in scope of Queen
			if column[c] || positiveDiag[r+c] || negativeDiag[r-c] {
				continue
			}

			// there is no Queen in scope so add the position in maps
			column[c] = true
			positiveDiag[r+c] = true
			negativeDiag[r-c] = true

			// prepare the string to append
			str := ""
			for i := 0; i < n; i++ {
				if i == c {
					str += "Q"
				} else {
					str += "."
				}
			}
			current = append(current, str)
			backtrack(r + 1)
			column[c] = false
			positiveDiag[r+c] = false
			negativeDiag[r-c] = false
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}
