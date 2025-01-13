package LinkedLists

func MergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	for len(lists) > 1 {
		mergedList := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}
			mergedList = append(mergedList, Merge(l1, l2))

		}
		lists = mergedList
	}
	return lists[0]
}

func Merge(list1, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	node := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	if list1 == nil {
		node.Next = list2
	}
	if list2 == nil {
		node.Next = list1
	}
	return head
}
