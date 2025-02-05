package Graphs

func ladderLength(beginWord string, endWord string, wordList []string) int {

	if !contains(wordList, endWord) {
		return 0
	}

	neighbors := make(map[string][]string)
	wordList = append(wordList, beginWord)

	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			neighbors[pattern] = append(neighbors[pattern], word)
		}
	}

	visited := make(map[string]bool)
	visited[beginWord] = true

	result := 1
	queue := []string{beginWord}

	for len(queue) > 0 {
		for i := len(queue); i > 0; i-- {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return result
			}

			for j := 0; j < len(word); j++ {
				pattern := word[:j] + "*" + word[j+1:]
				for _, neighborWord := range neighbors[pattern] {
					if !visited[neighborWord] {
						visited[neighborWord] = true
						queue = append(queue, neighborWord)
					}
				}
			}
		}
		result++
	}

	return 0
}

func contains(wordList []string, word string) bool {

	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}
