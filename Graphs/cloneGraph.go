package Graphs

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {

	if node == nil {
		return node
	}

	clones := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {

		if _, exists := clones[node]; exists {
			return clones[node]
		}

		copyNode := &Node{Val: node.Val}
		clones[node] = copyNode
		for _, neighbor := range node.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(neighbor))
		}
		return copyNode
	}

	return dfs(node)
}
