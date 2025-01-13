package LinkedLists

func FindDuplicate(nums []int) int {

	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		if slow == slow2 {
			return slow
		}
		slow = nums[slow]
		slow2 = nums[slow2]
	}
}
