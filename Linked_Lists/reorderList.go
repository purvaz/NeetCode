package LinkedLists

func ReorderList(head *ListNode) {

	ptr1 := head
	ptr2 := head.Next

	if ptr1.Next != nil {
		// Find the midpoint
		for ptr2 != nil && ptr2.Next != nil {
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next.Next
		}

		// Reverse the second half of the list
		current := ptr1.Next
		previous := &ListNode{}
		for current != nil {
			temp := current.Next
			if current == ptr1.Next {
				current.Next = nil
			} else {
				current.Next = previous
			}
			previous = current
			current = temp
		}
		ptr1.Next = nil
		ptr2 = previous
		ptr1 = head
		DisplayLinkedList(ptr1)
		DisplayLinkedList(ptr2)

		// rearrange
		temp := ptr1
		for ptr1 != nil && ptr2 != nil {
			temp = ptr1.Next
			ptr1.Next = ptr2
			ptr1 = temp
			temp = ptr2.Next
			ptr2.Next = ptr1
			ptr2 = temp
		}
		DisplayLinkedList(head)
	}

}
