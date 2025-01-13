package SlidingWindow

func MaxSlidingWindow(nums []int, k int) []int {

	maxArray := []int{}
	maximum := 0

	for i := 0; i < k; i++ {
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	maxArray = append(maxArray, maximum)

	l := 1
	r := k
	for r < len(nums) {
		if maximum == nums[l-1] {
			maximum = nums[l]
			for i := l; i <= r; i++ {
				if nums[i] > maximum {
					maximum = nums[i]
				}
			}
		}
		if maximum < nums[r] {
			maximum = nums[r]
		}
		maxArray = append(maxArray, maximum)
		l++
		r++
	}
	return maxArray
}
