package main

import "fmt"

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	// Trie node.
	// Each path represents a reversed suffix.
	type TrieNode struct {
		children [26]*TrieNode

		// bestIdx stores the index of the best wordContainer candidate
		// for the suffix represented by this node.
		//
		// "Best" means:
		// 1. shortest word length
		// 2. if same length, smaller original index
		bestIdx int
	}

	// Helper function to check whether candidate index cand is better
	// than current index curr.
	isBetter := func(cand, curr int) bool {
		// If there is no current best index yet, candidate is better.
		if curr == -1 {
			return true
		}

		// Prefer the shorter word.
		if len(wordsContainer[cand]) != len(wordsContainer[curr]) {
			return len(wordsContainer[cand]) < len(wordsContainer[curr])
		}

		// If lengths are the same, prefer the earlier index.
		return cand < curr
	}

	root := &TrieNode{
		bestIdx: -1,
	}

	// Build trie from wordsContainer.
	for i, word := range wordsContainer {
		node := root

		// Update root as well.
		// Root represents empty suffix.
		// If a query shares no suffix with any word,
		// answer should be the globally shortest word.
		if isBetter(i, node.bestIdx) {
			node.bestIdx = i
		}

		// Insert word in reversed order.
		// Example: "abcd" is inserted as 'd' -> 'c' -> 'b' -> 'a'.
		for j := len(word) - 1; j >= 0; j-- {
			ch := word[j] - 'a'

			if node.children[ch] == nil {
				node.children[ch] = &TrieNode{
					bestIdx: -1,
				}
			}

			node = node.children[ch]

			// At this node, this word has the suffix represented by the path.
			// Store the best index among all words sharing this suffix.
			if isBetter(i, node.bestIdx) {
				node.bestIdx = i
			}
		}
	}

	ans := make([]int, len(wordsQuery))

	// Answer each query.
	for i, query := range wordsQuery {
		node := root

		// Start with root.bestIdx.
		// This handles the case where no character matches at all.
		best := node.bestIdx

		// Walk query from back to front to match suffix characters.
		for j := len(query) - 1; j >= 0; j-- {
			ch := query[j] - 'a'

			// Cannot continue matching this suffix.
			if node.children[ch] == nil {
				break
			}

			node = node.children[ch]

			// This node represents a longer matched suffix,
			// so its bestIdx is the best answer for current matched suffix.
			best = node.bestIdx
		}

		ans[i] = best
	}

	return ans
}

func main() {
	fmt.Println(stringIndices([]string{"abcd", "bcd", "xbcd"}, []string{"cd", "bcd", "xyz"}))
}
