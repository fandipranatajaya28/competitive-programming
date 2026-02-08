package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j int, wordIdx int) bool {
	if wordIdx == len(word) {
		return true
	}

	// If out of boundary or char missmatch
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[wordIdx] {
		return false
	}

	// Change the current char so won't be visited again
	temp := board[i][j]
	board[i][j] = '*'

	// Check for 4 directions for the next match char
	isFound := dfs(board, word, i+1, j, wordIdx+1) ||
		dfs(board, word, i-1, j, wordIdx+1) ||
		dfs(board, word, i, j+1, wordIdx+1) ||
		dfs(board, word, i, j-1, wordIdx+1)

	// Return the board state
	board[i][j] = temp

	return isFound
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "SEE"))
}
