package Graphs

// Time complexity: O(V+E)

// Space complexity: O(V+E)

func validTree(n int, edges [][]int) bool {

	if n == 0 {
		return true
	}

	adjacencyList := make([][]int, n)
	for _, edge := range edges {
		v1, v2 := edge[0], edge[1]
		adjacencyList[v1] = append(adjacencyList[v1], v2)
		adjacencyList[v2] = append(adjacencyList[v2], v1)
	}

	visited := make(map[int]bool)

	var dfs func(current, parent int) bool
	dfs = func(current, parent int) bool {

		// Return false if already in visited set
		if visited[current] {
			return false
		}

		// perform dfs on adjacent nodes
		visited[current] = true
		for _, neighbor := range adjacencyList[current] {

			// if the neighbor is the parent, continue
			if neighbor == parent {
				continue
			}

			if !dfs(neighbor, current) {
				return false
			}
		}
		return true
	}

	return dfs(0, -1) && len(visited) == n
}
