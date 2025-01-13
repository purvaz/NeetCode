package Trees

import "fmt"

func RightSideView(root *TreeNode) []int {

	result := []int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelLength := len(queue)
		fmt.Println("Length of Queue : ", levelLength)
		rightMost := queue[levelLength-1]
		fmt.Println("Right most : ", rightMost.Val)
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node == rightMost {
				result = append(result, rightMost.Val)
			}
		}
	}
	return result
}

// func CreateBSTree() *TreeNode {
// 	node1 := &TreeNode{
// 		Val: 1,
// 	}
// 	node2 := &TreeNode{
// 		Val: 2,
// 	}
// 	node3 := &TreeNode{
// 		Val: 3,
// 	}
// 	node4 := &TreeNode{
// 		Val: 4,
// 	}
// 	node5 := &TreeNode{
// 		Val: 5,
// 	}

// 	node1.Left = node2
// 	node1.Right = node3

// 	node2.Right = node5

// 	node3.Right = node4
// 	return node1
// }
