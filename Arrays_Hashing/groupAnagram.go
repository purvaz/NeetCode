// Medium

package Arrayshashing

func GroupAnagrams(strs []string) [][]string {

	var anaMap = make(map[[26]int][]string)

	for _, str := range strs {
		charArr := [26]int{0}
		for i := 0; i < len(str); i++ {
			charArr[str[i]-'a']++
		}
		if _, exists := anaMap[charArr]; exists {
			anaMap[charArr] = append(anaMap[charArr], str)
		} else {
			anaMap[charArr] = append(anaMap[charArr], str)
		}
	}

	result := [][]string{}
	for _, val := range anaMap {
		result = append(result, val)
	}
	return result
}
