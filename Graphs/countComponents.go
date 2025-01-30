package Graphs

type DSU struct {
	parent []int
	rank   []int
}

// Create a struct that will have the parent and rank array.
func NewDSU(n int) *DSU {

	dsu := &DSU{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	// Populate the parent and rank arrays.
	// The parent array will hold the parent to the current node,
	// whereas the rank contains the number of children that the current node has.
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = 1
	}
	return dsu
}

func (dsu *DSU) Find(node int) int {

	current := node
	// Check if the parent of current node is the node itself. If not, iterate till we find such a node.
	for current != dsu.parent[current] {
		// Assign the root parent as the parent of the node instead of the immediate parent.
		dsu.parent[current] = dsu.parent[dsu.parent[current]]
		current = dsu.parent[current]
	}
	return current
}

func (dsu *DSU) Union(node1, node2 int) bool {

	parent1 := dsu.Find(node1)
	parent2 := dsu.Find(node2)

	// If the parents are same
	if parent1 == parent2 {
		return false
	}
	// Check the parent hierarchy and assign the rank and the parent to the nodes.
	if dsu.rank[parent2] > dsu.rank[parent1] {
		parent1, parent2 = parent2, parent1
	}
	dsu.parent[parent2] = parent1
	dsu.rank[parent1] += dsu.rank[parent2]

	return true
}

func countComponents(n int, edges [][]int) int {

	dsu := NewDSU(n)

	result := n

	// For every edge, call the union function
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		// If the union function returns a true, decrement the result.
		// This means that the nodes are connected.
		if dsu.Union(u, v) {
			result--
		}
	}
	return result
}
