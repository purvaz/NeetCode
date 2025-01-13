package TwoPointers

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {

	numMap := make(map[[3]int]bool)
	result := [][]int{}
	// sort the array first
	sort.Ints(nums)

	for start, mid, end := 0, 1, len(nums)-1; start < end; {
		for mid < end {
			if nums[start]+nums[mid]+nums[end] == 0 {
				numMap[[3]int{nums[start], nums[mid], nums[end]}] = true
				mid++
				end--
			} else {
				if nums[start]+nums[mid]+nums[end] < 0 {
					mid++
				} else {
					end--
				}
			}
		}
		if mid >= end {
			start++
			mid = start + 1
			end = len(nums) - 1
		}
	}
	for key := range numMap {
		result = append(result, append([]int{}, key[:]...))
	}
	return result
}
