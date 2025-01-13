package Trees

var limit int

func kthSmallest(root *TreeNode, k int) int {

	stack := make([]*TreeNode, 0, k)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
	return -1
}

func CreateBSTree() *TreeNode {
	node1 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 1,
	}
	node3 := &TreeNode{
		Val: 4,
	}
	node4 := &TreeNode{
		Val: 2,
	}

	node1.Left = node2
	node1.Right = node3

	node2.Right = node4

	return node1
}
