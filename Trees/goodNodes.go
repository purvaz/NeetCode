package Trees

func GoodNodes(root *TreeNode) int {

	return dfsMax(root, root.Val)
}

func dfsMax(root *TreeNode, max int) int {

	result := 0

	if root == nil {
		return 0
	}

	if root.Val >= max {
		max = root.Val
		result++
	}
	result += dfsMax(root.Left, max)
	result += dfsMax(root.Right, max)

	return result
}
