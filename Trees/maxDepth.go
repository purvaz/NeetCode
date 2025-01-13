package Trees

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	return 1 + maximum(maxDepth(root.Left), maxDepth(root.Right))
}

// func maximum(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
