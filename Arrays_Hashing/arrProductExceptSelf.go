package Arrayshashing

import "fmt"

func ProductExceptSelf(nums []int) []int {

	prodArr := make([]int, len(nums))

	prefix := 1
	for key, val := range nums {
		prodArr[key] = prefix
		prefix = prefix * val
	}
	fmt.Println(prefix, prodArr)

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		prodArr[i] *= postfix
		postfix *= nums[i]
		fmt.Println(postfix, prodArr)
	}
	return prodArr
}
