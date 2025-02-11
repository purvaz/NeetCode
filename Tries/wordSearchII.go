package Tries

type Node struct {
	children    map[byte]*Node
	isEndOfWord bool
}

func (this *Node) addWords(word string) {

	current := this

	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := current.children[char]; !exists {
			current.children[char] = &Node{children: map[byte]*Node{}}
		}
		current = current.children[char]
	}
	current.isEndOfWord = true
}

func findWords(board [][]byte, words []string) []string {

	root := &Node{children: make(map[byte]*Node)}

	for _, word := range words {
		root.addWords(word)
	}

	rows, cols := len(board), len(board[0])
	var result []string
	visited := make(map[[2]int]bool)
	wordSet := make(map[string]bool)

	var dfs func(r, c int, node *Node, word string)
	dfs = func(r, c int, node *Node, word string) {

		if r < 0 || c < 0 || r == rows || c == cols || visited[[2]int{r, c}] || board[r][c] == 0 {
			return
		}

		char := board[r][c]

		nextNode, isPresent := node.children[char]
		if !isPresent {
			return
		}

		visited[[2]int{r, c}] = true
		word += string(char)
		if nextNode.isEndOfWord {
			wordSet[word] = true
		}

		dfs(r+1, c, nextNode, word)
		dfs(r-1, c, nextNode, word)
		dfs(r, c+1, nextNode, word)
		dfs(r, c-1, nextNode, word)

		visited[[2]int{r, c}] = false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root, "")
		}
	}

	for word := range wordSet {
		result = append(result, word)
	}

	return result

}
