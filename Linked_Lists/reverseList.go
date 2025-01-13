package LinkedLists

func ReverseList(head *ListNode) *ListNode {

	current := head
	previous := &ListNode{}

	if head == nil {
		return nil
	}
	for current != nil {
		temp := current.Next
		if current == head {
			current.Next = nil
		} else {
			current.Next = previous
		}
		previous = current
		current = temp
	}

	return previous
}
