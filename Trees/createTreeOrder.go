package Trees

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var mid int
	root := &TreeNode{Val: preorder[0]}

	for index, value := range inorder {
		if value == preorder[0] {
			mid = index
			break
		}
	}
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
