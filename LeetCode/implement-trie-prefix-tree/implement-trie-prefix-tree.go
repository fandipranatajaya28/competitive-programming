package main

import "fmt"

type Trie struct {
	isLeaf bool
	child  map[byte]*Trie
}

func Constructor() Trie {
	object := new(Trie)
	object.child = make(map[byte]*Trie)
	return *object
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := range len(word) {
		if _, isExist := curr.child[word[i]]; !isExist {
			child := Constructor()
			curr.child[word[i]] = &child
		}
		curr = curr.child[word[i]]
	}
	curr.isLeaf = true
	return
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := range len(word) {
		if _, isExist := curr.child[word[i]]; !isExist {
			return false
		}
		curr = curr.child[word[i]]
	}
	return curr.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := range len(prefix) {
		if _, isExist := curr.child[prefix[i]]; !isExist {
			return false
		}
		curr = curr.child[prefix[i]]
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
