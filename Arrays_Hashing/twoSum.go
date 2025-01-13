// Easy

package Arrayshashing

func TwoSum(nums []int, target int) []int {

	var numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exists := numMap[target-nums[i]]; exists {
			return []int{numMap[target-nums[i]], i}
		} else {
			numMap[nums[i]] = i
		}
	}
	return nil
}
