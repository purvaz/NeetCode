package BinarySearch

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// arr1 will always be longer than arr2
	arr1, arr2 := nums1, nums2
	total := len(arr1) + len(arr2)
	half := (total + 1) / 2

	if len(nums1) > len(nums2) {
		arr1, arr2 = arr2, arr1
	}

	var left1, left2, right1, right2 float64
	l, r := 0, len(arr1)-1

	for {
		i := (l + r) >> 1
		j := half - i - 2

		if i >= 0 {
			left1 = float64(arr1[i])
		} else {
			left1 = math.Inf(-1)
		}

		if (i + 1) < len(arr1) {
			right1 = float64(arr1[i+1])
		} else {
			right1 = math.Inf(1)
		}

		if j >= 0 {
			left2 = float64(arr2[j])
		} else {
			left2 = math.Inf(-1)
		}

		if (j + 1) < len(arr2) {
			right2 = float64(arr2[j+1])
		} else {
			right2 = math.Inf(1)
		}

		if left1 <= right2 && left2 <= right1 {
			if total%2 == 1 {
				return maximum(left1, left2)
			}
			return (maximum(left1, left2) + minimum(right1, right2)) / 2
		} else if left1 > right2 {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func minimum(num1, num2 float64) float64 {

	if num1 > num2 {
		return num2
	}
	return num1
}

func maximum(num1, num2 float64) float64 {

	if num1 < num2 {
		return num2
	}
	return num1
}
