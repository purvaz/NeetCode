package SlidingWindow

func CharacterReplacement(s string, k int) int {

	maxLength := 0
	var charMap = make(map[byte]int)
	l := 0
	for r, _ := range s {
		charMap[s[r]] = 1 + charMap[s[r]]
		if (r-l+1)-getFrequentCharacter(charMap) > k {
			charMap[s[l]]--
			l++
		}
		if maxLength < (r - l + 1) {
			maxLength = (r - l + 1)
		}
	}
	return maxLength
}

func getFrequentCharacter(charMap map[byte]int) int {
	frequent := 0
	for _, v := range charMap {
		if v > frequent {
			frequent = v
		}
	}
	return frequent
}
