package LinkedLists

func ReverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{Val: 0, Next: head}
	groupPrev := dummy

	for {
		kth := getKth(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next

		// Reverse
		prev, curr := kth.Next, groupPrev.Next
		for curr != groupNext {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		temp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = temp

	}
	DisplayLinkedList(dummy.Next)
	return dummy.Next
}

func getKth(current *ListNode, k int) *ListNode {
	for current != nil && k > 0 {
		current = current.Next
		k--
	}
	if k > 0 {
		return nil
	}
	return current
}
