package Stack

func IsValid(s string) bool {

	paranMap := map[byte]byte{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	paranStack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			paranStack = append(paranStack, s[i])
		} else {
			if len(paranStack) > 0 && paranMap[paranStack[len(paranStack)-1]] == s[i] {
				paranStack = paranStack[:len(paranStack)-1]
			} else {
				return false
			}
		}
	}
	return len(paranStack) == 0
}
