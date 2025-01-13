package TwoPointers

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {

	str := regexp.MustCompile("[a-zA-z0-9]").FindAllString(strings.ToLower(s), -1)
	i, j := 0, len(str)-1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
