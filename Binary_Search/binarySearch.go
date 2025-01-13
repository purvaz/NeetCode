package BinarySearch

func BinarySearching(nums []int, target int) int {

	start, end := 0, len(nums)-1
	for start <= end {
		mid := (end + start) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
