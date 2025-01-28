package Graphs

func findOrder(numCourses int, prerequisites [][]int) []int {

	prereqMap := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		prereqMap[i] = []int{}
	}

	for _, prerequisite := range prerequisites {
		course, prereq := prerequisite[0], prerequisite[1]
		prereqMap[course] = append(prereqMap[course], prereq)
	}

	output := []int{}
	visit := make(map[int]bool)
	cycle := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(crs int) bool {
		if cycle[crs] {
			return false
		}
		if visit[crs] {
			return true
		}
		cycle[crs] = true

		for _, pre := range prereqMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		cycle[crs] = false
		visit[crs] = true
		output = append(output, crs)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return output
}
