package Trees

var result int

func maxPathSum(root *TreeNode) int {
	result = root.Val
	return DFS(root)
}

func DFS(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// get the max of left subtree
	leftMax := DFS(root.Left)

	// handling negative values
	leftMax = max(leftMax, 0)

	// get the max of right subtree
	rightMax := DFS(root.Right)

	// handling negative values
	rightMax = max(rightMax, 0)

	// calculating the max path WITH a split
	result = max(result, (leftMax + root.Val + rightMax))

	// return the path value WITHOUT a split
	return root.Val + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
