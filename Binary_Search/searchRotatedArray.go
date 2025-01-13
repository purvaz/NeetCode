package BinarySearch

func Search(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
