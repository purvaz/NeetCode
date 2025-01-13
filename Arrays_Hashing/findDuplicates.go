package Arrayshashing

import "math"

func FindDuplicates(nums []int) []int {

	result := []int{}

	for i := 0; i < len(nums); i++ {

		// getting the value at i
		num := int(math.Abs(float64(nums[i])))

		if nums[num-1] < 0 {
			result = append(result, num)
		}
		nums[num-1] *= -1

	}
	return result
}
