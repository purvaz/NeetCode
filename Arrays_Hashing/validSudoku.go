package Arrayshashing

func IsValidSudoku(board [][]byte) bool {

	var hashMap = make(map[string]bool)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			current_value := string(board[r][c])
			if current_value == "." {
				continue
			}
			_, okr := hashMap[current_value+" in row "+string(r)]
			_, okc := hashMap[current_value+" in col "+string(c)]
			_, oks := hashMap[current_value+" in square "+string(r/3)+string(c/3)]

			if okr || okc || oks {
				return false
			} else {
				hashMap[current_value+" in row "+string(r)] = true
				hashMap[current_value+" in col "+string(c)] = true
				hashMap[current_value+" in square "+string(r/3)+string(c/3)] = true
			}
		}
	}
	return false
}
