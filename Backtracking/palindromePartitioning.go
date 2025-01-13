package Backtracking

import "fmt"

func Partition(s string) [][]string {

	result := make([][]string, 0)
	subset := make([]string, 0)

	var dfs func(index int)
	dfs = func(index int) {

		if index == len(s) {
			result = append(result, append([]string{}, subset...))
			return
		}

		for j := index; j < len(s); j++ {
			fmt.Println(s[j])
			if isPalindrome(s, index, j) {
				subset = append(subset, s[index:j+1])
				dfs(j + 1)
				subset = subset[:len(subset)-1]
			}
		}
	}
	dfs(0)
	return result
}

func isPalindrome(s string, l, r int) bool {

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
