package Trees

func LevelOrder(root *TreeNode) [][]int {

	result := [][]int{}

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		lenQueue := len(queue)
		for i := 0; i < lenQueue; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				level = append(level, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		result = append(result, level)
	}
	return result
}

func CreateTree() *TreeNode {
	node1 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 9,
	}
	node3 := &TreeNode{
		Val: 20,
	}
	node4 := &TreeNode{
		Val: 15,
	}
	node5 := &TreeNode{
		Val: 7,
	}

	node1.Left = node2
	node1.Right = node3

	node3.Left = node4
	node3.Right = node5

	return node1
}
