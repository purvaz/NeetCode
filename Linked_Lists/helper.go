package LinkedLists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {

	curr := &ListNode{}
	isCreated := false
	for i := len(nums) - 1; i >= 0; i-- {
		prev := &ListNode{}
		curr.Val = nums[i]
		if !isCreated {
			curr.Next = nil
			isCreated = true
		}
		prev.Next = curr
		if i > 0 {
			curr = prev
		}
	}
	return curr
}

func DisplayLinkedList(head *ListNode) {
	fmt.Println("------------------")
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println("\n------------------")
}
