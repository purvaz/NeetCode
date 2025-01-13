// Easy

package Arrayshashing

import "fmt"

func IsAnagram(s string, t string) bool {

	var charArr = [26]int{0}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		charArr[int(s[i])-'a']++
		charArr[int(t[i])-'a']--
		fmt.Println(charArr)
	}

	for i := 0; i < len(charArr); i++ {
		if charArr[i] != 0 {
			return false
		}
	}

	return true

}
