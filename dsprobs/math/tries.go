package math

import "fmt"

// Issues while coding and testing
//   - alphabet count Not set,
//     temp.children[charIndex] != nil should be temp.children[charIndex] == nil
const AlphabetCount = 26

type TrieNode struct {
	isEnd    bool
	children [AlphabetCount]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) TrieInsert(w string) {
	temp := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			temp.children[charIndex] = &TrieNode{}
		}
		temp = temp.children[charIndex]
	}
	temp.isEnd = true

}

func (t *Trie) TrieSearch(w string) bool {
	temp := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			return false
		}
		temp = temp.children[charIndex]
	}
	return temp.isEnd
}

func initTrie() *Trie {
	myTrie := &Trie{root: &TrieNode{}}
	fmt.Print(myTrie)
	return myTrie
}
