package Trees

import "math"

type BalancedTree struct {
	Balanced bool
	Height   int
}

func isBalanced(root *TreeNode) bool {

	return traverse(root).Balanced
}

func traverse(root *TreeNode) BalancedTree {

	if root == nil {
		return BalancedTree{
			Balanced: true,
			Height:   0,
		}
	}

	left := traverse(root.Left)
	right := traverse(root.Right)

	balanced := (left.Balanced && right.Balanced && (math.Abs(float64(left.Height)-float64(right.Height))) <= 1)

	return BalancedTree{
		Balanced: balanced,
		Height:   1 + maximum(left.Height, right.Height),
	}
}
