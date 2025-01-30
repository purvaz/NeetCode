package Graphs

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {

	dsu := &DSU{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = 1
	}
	return dsu
}

func (dsu *DSU) Find(node int) int {

	current := node
	for current != dsu.parent[current] {
		dsu.parent[current] = dsu.parent[dsu.parent[current]]
		current = dsu.parent[current]
	}
	return current
}

func (dsu *DSU) Union(node1, node2 int) bool {

	parent1 := dsu.Find(node1)
	parent2 := dsu.Find(node2)

	if parent1 == parent2 {
		return false
	}
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

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if dsu.Union(u, v) {
			result--
		}
	}
	return result
}
