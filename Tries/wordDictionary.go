package Tries

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

type WordTrieNode struct {
	children [26]*WordTrieNode
	wordEnd  bool
}

type WordDictionary struct {
	root *WordTrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &WordTrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {

	current := this.root
	for _, char := range word {
		index := char - 'a'
		if current.children[index] == nil {
			current.children[index] = &WordTrieNode{}
		}
		current = current.children[index]
	}
	current.wordEnd = true
}

func (this *WordDictionary) Search(word string) bool {

}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool {

	current := root

	for i := j; i < len(word); i++ {
		char := word[i]

		if char == '.' {
			for _, child := range current.children {
				if child != nil && this.dfs(word, i+1, child) {
					return true
				}
			}
			return false
		} else {
			index := char - 'a'
			if current.children[rune(index)] == nil {
				return false
			}
			current = current.children[rune(index)]
		}
	}
	return current.endOfWord
}
