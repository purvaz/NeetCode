package Trees

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

// func CreateTree() *TreeNode {
// 	node1 := &TreeNode{
// 		Val: 20,
// 	}
// 	node2 := &TreeNode{
// 		Val: 10,
// 	}
// 	node3 := &TreeNode{
// 		Val: 30,
// 	}
// 	node4 := &TreeNode{
// 		Val: 25,
// 	}
// 	node5 := &TreeNode{
// 		Val: 40,
// 	}
// 	node6 := &TreeNode{
// 		Val: 22,
// 	}

// 	node1.Left = node2
// 	node3.Left = node4
// 	node4.Left = node6

// 	node1.Right = node3
// 	node3.Right = node5

// 	return node1
// }

func TraverseTree(root *TreeNode) {

	preOrder(root)
	fmt.Println("")
	inOrder(root)
	fmt.Println("")
	postOrder(root)
}

func preOrder(root *TreeNode) {

	if root == nil {
		return
	}
	fmt.Print(root.Val, " --> ")
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *TreeNode) {

	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Print(root.Val, " --> ")
	inOrder(root.Right)
}

func postOrder(root *TreeNode) {

	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Print(root.Val, " --> ")
}
