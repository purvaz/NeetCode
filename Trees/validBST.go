package Trees

func isValidBST(root *TreeNode) bool {
	return checkValidity(root, nil, nil)
}

func checkValidity(root, min, max *TreeNode) bool {

	if root == nil {
		return true
	}
	if min != nil && root.Val < min.Val {
		return false
	}
	if max != nil && root.Val > max.Val {
		return false
	}

	return checkValidity(root.Left, min, root) && checkValidity(root.Right, root, max)
}
