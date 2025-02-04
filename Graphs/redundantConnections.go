package Graphs

func findRedundantConnection(edges [][]int) []int {

	n := len(edges)
	adjacencyList := make([][]int, n+1)
	visited := make([]bool, n+1)

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {

		if visited[node] {
			return true
		}

		visited[node] = true
		for _, neighbor := range adjacencyList[node] {
			if neighbor == parent {
				continue
			}

			if dfs(neighbor, node) {
				return true
			}
		}

		return false
	}

	for _, edge := range edges {

		u, v := edge[0], edge[1]
		adjacencyList[u] = append(adjacencyList[u], v)
		adjacencyList[v] = append(adjacencyList[v], u)

		for i := 0; i <= n; i++ {
			visited[i] = false
		}

		if dfs(u, -1) {
			return []int{u, v}
		}
	}

	return []int{}
}
