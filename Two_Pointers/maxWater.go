package TwoPointers

func getMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func MaxArea(height []int) int {

	maximum := 0
	for i, j := 0, len(height)-1; i < j; {
		minimum := getMin(height[i], height[j])
		// if height[i] < height[j] {
		// 	minimum = height[i]
		// } else {
		// 	minimum = height[j]
		// }
		area := (j - i) * minimum
		if maximum < area {
			maximum = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maximum
}
