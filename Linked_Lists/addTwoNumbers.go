package LinkedLists

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	var head *ListNode
	prev := head

	for l1 != nil || l2 != nil {
		temp := &ListNode{}

		// Gather values
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		// Calculate sum and carry
		sum := v1 + v2 + carry
		carry = sum / 10
		temp.Val = sum % 10

		// Create and link the list
		if head == nil {
			head = temp
		} else {
			prev.Next = temp
		}
		prev = temp

		// Go to next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// Handle edge case of carry at the end
	if carry != 0 {
		temp := &ListNode{Val: carry, Next: nil}
		prev.Next = temp
	}
	return head
}
