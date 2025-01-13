package TwoPointers

func TwoSum(numbers []int, target int) []int {

	start, end := 0, len(numbers)-1
	for start < end {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		}
		if numbers[start]+numbers[end] > target {
			end--
		} else {
			start++
		}
	}
	return []int{}
}
