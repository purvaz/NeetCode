package DynamicProgramming

func CanPartition(nums []int) bool {

	sumSet := make(map[int]bool)
	sum := getSum(nums)
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	sumSet[0] = true
	sumSet[nums[len(nums)-1]] = true

	for i := len(nums) - 2; i >= 0; i-- {
		temp := make(map[int]bool)
		for key := range sumSet {
			if key+nums[i] == target {
				return true
			}
			temp[key+nums[i]] = true
		}
		for key := range temp {
			sumSet[key] = true
		}
	}
	return false
}

func getSum(nums []int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
