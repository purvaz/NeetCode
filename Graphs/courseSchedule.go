package Graphs

func canFinish(numCourses int, prerequisites [][]int) bool {

	prereqMap := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		prereqMap[i] = []int{}
	}

	for _, prerequisite := range prerequisites {
		course, prereq := prerequisite[0], prerequisite[1]
		prereqMap[course] = append(prereqMap[course], prereq)
	}

	visited := make(map[int]bool)

	var dfs func(course int) bool
	dfs = func(course int) bool {

		if visited[course] {
			return false
		}

		if len(prereqMap[course]) == 0 {
			return true
		}

		visited[course] = true

		for _, pre := range prereqMap[course] {
			if !dfs(pre) {
				return false
			}
		}

		visited[course] = false
		prereqMap[course] = []int{}

		return true
	}

	for c := 0; c < numCourses; c++ {
		if !dfs(c) {
			return false
		}
	}
	return true
}
