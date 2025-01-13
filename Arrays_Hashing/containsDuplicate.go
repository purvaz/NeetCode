// Easy

package Arrayshashing

func ContainsDuplicate(nums []int) bool {

	var hashMap = make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, exists := hashMap[nums[i]]; exists {
			return true
		} else {
			hashMap[nums[i]] = true
		}
	}
	return false
}
