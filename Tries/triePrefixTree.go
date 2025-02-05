package Tries

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {

	current := this.root

	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[char]
	}
	current.endOfWord = true
}

func (this *Trie) Search(word string) bool {

	current := this.root
	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return current.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {

	current := this.root
	for _, char := range prefix {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
