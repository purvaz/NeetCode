package SlidingWindow

func CheckInclusion(s1 string, s2 string) bool {

	s1Hashmap := [26]int{}
	s2Hashmap := [26]int{}

	if len(s2) < len(s1) {
		return false
	}
	// Populating map for s1
	for i := 0; i < len(s1); i++ {
		s1Hashmap[s1[i]-'a']++
	}

	l := 0
	r := len(s1) - 1
	for i := 0; i < len(s1); i++ {
		s2Hashmap[s2[i]-'a']++
	}
	for r < len(s2)-1 {
		if s1Hashmap == s2Hashmap {
			return true
		} else {
			s2Hashmap[s2[l]-'a']--
			l++
			r++
			s2Hashmap[s2[r]-'a']++
		}

	}
	return s1Hashmap == s2Hashmap
}
