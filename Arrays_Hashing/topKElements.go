package Arrayshashing

import "fmt"

func TopKFrequent(nums []int, k int) []int {

	var kMap = make(map[int]int)
	var countArr = make([][]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		if _, exists := kMap[nums[i]]; exists {
			kMap[nums[i]]++
		} else {
			kMap[nums[i]] = 1
		}
	}

	for num, count := range kMap {
		countArr[count] = append(countArr[count], num)
	}
	fmt.Println(countArr)
	res := []int{}

	for i := len(countArr) - 1; i > 0; i-- {
		res = append(res, countArr[i]...)
		if len(res) == k {
			return res
		}
	}
	return []int{}
}
