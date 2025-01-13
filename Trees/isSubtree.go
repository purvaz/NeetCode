package Trees

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}

	return (equals(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
}

func equals(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil || root.Val != subRoot.Val {
		return false
	}
	return equals(root.Left, subRoot.Left) && equals(root.Right, subRoot.Right)
}
