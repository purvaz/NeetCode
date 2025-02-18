package DynamicProgramming2D

func isInterleave(s1 string, s2 string, s3 string) bool {

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(s3) {
			return i == len(s1) && j == len(s2)
		}

		if i < len(s1) && s1[i] == s3[k] {
			if dfs(i+1, j, k+1) {
				return true
			}
		}

		if j < len(s2) && s2[j] == s3[k] {
			if dfs(i, j+1, k+1) {
				return true
			}
		}

		return false
	}

	return dfs(0, 0, 0)

}
