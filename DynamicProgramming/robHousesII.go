package DynamicProgramming

func robII(nums []int) int {

	return maximum(nums[0], maximum(robberHelper(nums[1:]), robberHelper(nums[:len(nums)-1])))
}

func robberHelper(nums []int) int {

	rob1, rob2 := 0, 0

	for _, num := range nums {
		temp := maximum(rob1+num, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func maximum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
