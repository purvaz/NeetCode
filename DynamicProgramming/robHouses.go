package DynamicProgramming

func rob(nums []int) int {

	house1, house2 := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		temp := max((house1 + nums[i]), house2)
		house1 = house2
		house2 = temp
	}
	return house2
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
