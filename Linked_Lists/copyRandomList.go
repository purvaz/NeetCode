package LinkedLists

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CreateRandomList() *Node {

	node1 := &Node{
		Val: 7,
	}

	node2 := &Node{
		Val: 13,
	}

	node3 := &Node{
		Val: 11,
	}

	node4 := &Node{
		Val: 10,
	}

	node5 := &Node{
		Val: 1,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = nil

	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	return node1
}

func CopyRandomList(head *Node) *Node {

	copyMap := make(map[*Node]*Node)

	current := head
	for current != nil {
		temp := &Node{Val: current.Val}
		copyMap[current] = temp
		current = current.Next
	}

	current = head
	for current != nil {
		temp := copyMap[current]
		temp.Next = copyMap[current.Next]
		temp.Random = copyMap[current.Random]
		current = current.Next
	}
	return copyMap[head]
}
