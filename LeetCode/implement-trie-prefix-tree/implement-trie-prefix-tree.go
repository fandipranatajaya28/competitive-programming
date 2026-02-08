package main

import "fmt"

type Trie struct {
	child map[byte]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		child: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if curr.child[char] == nil {
			child := Constructor()
			curr.child[char] = &child
		}
		curr = curr.child[char]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := range len(word) {
		char := word[i]
		if curr.child[char] == nil {
			return false
		}
		curr = curr.child[char]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := range len(prefix) {
		char := prefix[i]
		if curr.child[char] == nil {
			return false
		}
		curr = curr.child[char]
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

func main() {
	var (
		commands = []string{"insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search"}
		inputs   = []string{"app", "apple", "beer", "add", "jam", "rental", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam"}
		trie     = Constructor()
	)

	for idx, command := range commands {
		switch command {
		case "insert":
			trie.Insert(inputs[idx])
		case "search":
			fmt.Println(trie.Search(inputs[idx]), inputs[idx])
		case "startsWith":
			fmt.Println(trie.StartsWith(inputs[idx]), inputs[idx])
		}
	}
}
