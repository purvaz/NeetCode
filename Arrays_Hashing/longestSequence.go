package Arrayshashing

func LongestConsecutive(nums []int) int {

	elementMap := make(map[int]bool)

	for _, value := range nums {
		elementMap[value] = true
	}

	maximum := 0
	for i := 0; i < len(nums); i++ {
		if !elementMap[nums[i]-1] {
			length := 0
			for elementMap[nums[i]+length] {
				length++
			}
			if maximum < length {
				maximum = length
			}
		}
	}
	return maximum
}
