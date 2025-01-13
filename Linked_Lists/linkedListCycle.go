package LinkedLists

func HasCycle(head *ListNode) bool {

	fast, slow := head, head

	if head == nil {
		return false
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}
