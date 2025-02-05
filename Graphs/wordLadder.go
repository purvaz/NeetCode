package Graphs

func ladderLength(beginWord string, endWord string, wordList []string) int {

	// check if the end word is in the word list
	if !contains(wordList, endWord) {
		return 0
	}

	// make an adjacency list that will store all the patterns and the words that match the pattern
	neighbors := make(map[string][]string)
	// append the begin word to the wordlist
	wordList = append(wordList, beginWord)

	// traverse through the word list, and create the patterns for that word
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			neighbors[pattern] = append(neighbors[pattern], word)
		}
	}

	// create a visited map that will check if the particular word has been visited
	visited := make(map[string]bool)
	visited[beginWord] = true

	result := 1
	// add the first word in the queue
	queue := []string{beginWord}

	// perform BFS on the queue
	for len(queue) > 0 {
		for i := len(queue); i > 0; i-- {
			word := queue[0]
			queue = queue[1:]

			// check if the word matches the end word
			if word == endWord {
				return result
			}

			for j := 0; j < len(word); j++ {
				// create the patterns for the particular word,
				pattern := word[:j] + "*" + word[j+1:]
				for _, neighborWord := range neighbors[pattern] {
					if !visited[neighborWord] {
						visited[neighborWord] = true
						// add the words that match the pattern in the queue
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
