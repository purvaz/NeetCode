package BinarySearch

import "fmt"

func FindMin(nums []int) int {

	left, right := 0, len(nums)-1

	minimum := 5001
	for left <= right {
		mid := (right + left) / 2
		fmt.Println("Mid is : ", nums[mid])
		// Already sorted, so return leftmost element
		if nums[right] > nums[left] {
			if nums[left] < minimum {
				fmt.Println("Array is sorted, returning min. ", nums[left])
				minimum = nums[left]
				break
			}
		}
		if nums[mid] < minimum {
			// check if mid is lesser than min
			minimum = nums[mid]
			fmt.Println("Comparing with min ", nums[mid])
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return minimum
}
