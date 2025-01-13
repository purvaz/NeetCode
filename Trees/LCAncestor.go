package Trees

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	current := root
	if p.Val < current.Val && q.Val < current.Val {
		current = current.Left
	} else if p.Val > current.Val && q.Val > current.Val {
		current = current.Right
	} else if (p.Val < current.Val && q.Val > current.Val) || p.Val == current.Val || q.Val == p.Val {
		return current
	}
	return nil
}

// func CreateBSTree() *TreeNode {
// 	node1 := &TreeNode{
// 		Val: 6,
// 	}
// 	node2 := &TreeNode{
// 		Val: 2,
// 	}
// 	node3 := &TreeNode{
// 		Val: 8,
// 	}
// 	node4 := &TreeNode{
// 		Val: 0,
// 	}
// 	node5 := &TreeNode{
// 		Val: 4,
// 	}
// 	node6 := &TreeNode{
// 		Val: 7,
// 	}
// 	node7 := &TreeNode{
// 		Val: 9,
// 	}
// 	node8 := &TreeNode{
// 		Val: 3,
// 	}
// 	node9 := &TreeNode{
// 		Val: 5,
// 	}

// 	node1.Left = node2
// 	node1.Right = node3

// 	node2.Left = node4
// 	node2.Right = node5

// 	node3.Left = node6
// 	node3.Right = node7

// 	node5.Left = node8
// 	node5.Right = node9

// 	return node1
// }
