package LinkedLists

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	l := head
	r := head
	for n > 0 && r != nil {
		r = r.Next
		n--
	}

	if r == nil {
		return l.Next
	}
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return head
}
