package Trees

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	result := make([]string, 0)

	// create a function for DFS
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		// if the node is null, append a "N" in the string
		if node == nil {
			result = append(result, "N")
			return
		}
		// else, append the value at the node in the string
		result = append(result, strconv.Itoa(node.Val))
		// call the preorder function recursively, passing the left and right nodes
		preorder(node.Left)
		preorder(node.Right)
	}

	// calling the function, passing the root node
	preorder(root)
	// finally, return the result string
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	// create an array called "values" that will store all the values of nodes
	values := strings.Split(data, ",")

	// maintain i to keep track of the # of values
	i := 0

	// creating a dfs function
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		// if the value at i is "N", it means the node is nil
		if values[i] == "N" {
			i++
			return nil
		}
		// else, convert the string to int, and create an empty node and assign the value
		value, _ := strconv.Atoi(values[i])
		node := &TreeNode{Val: value}
		i++

		// recursively call the dfs function on left and right
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
