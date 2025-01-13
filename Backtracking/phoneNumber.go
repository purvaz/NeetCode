package Backtracking

func LetterCombinations(digits string) []string {

	result := make([]string, 0)

	if len(digits) == 0 {
		return result
	}

	digitMap := make(map[byte]string)
	digitMap['2'] = "abc"
	digitMap['3'] = "def"
	digitMap['4'] = "ghi"
	digitMap['5'] = "jkl"
	digitMap['6'] = "mno"
	digitMap['7'] = "pqrs"
	digitMap['8'] = "tuv"
	digitMap['9'] = "wxyz"

	var backtrack func(index int, substr string)
	backtrack = func(index int, substr string) {
		if len(substr) == len(digits) {
			result = append(result, substr)
			return
		}

		chars := digitMap[digits[index]]
		for i := 0; i < len(chars); i++ {
			substr += string(chars[i])
			backtrack(index+1, substr)
			substr = substr[:len(substr)-1]
		}
	}

	backtrack(0, "")
	return result
}
