package main

import "fmt"

func solve(board [][]byte) {
	if len(board) < 2 || len(board[0]) < 2 {
		return
	}

	// Mark cells connected to border as border_connected('B')
	for i := 0; i < len(board); i++ {
		colorUsingDFS(board, i, 0)               // traverse on first column
		colorUsingDFS(board, i, len(board[0])-1) // traverse on last column
	}
	for j := 0; j < len(board[0]); j++ {
		colorUsingDFS(board, 0, j)            // traverse on first row
		colorUsingDFS(board, len(board)-1, j) // traverse on last row
	}

	// iterate over all cells and mark all 'B' as 'O' and rest as 'X'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'B' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
			fmt.Println(string(board[i][j]))
		}
	}
}

func colorUsingDFS(board [][]byte, i, j int) {
	// If current border connected element is not 'O', skip it
	if board[i][j] != 'O' {
		return
	}

	// Make current border connected element as 'B'
	board[i][j] = 'B'

	// Check top cell
	if i-1 >= 0 {
		colorUsingDFS(board, i-1, j)
	}

	// Check bottom cell
	if i+1 < len(board) {
		colorUsingDFS(board, i+1, j)
	}

	// Check left cell
	if j-1 >= 0 {
		colorUsingDFS(board, i, j-1)
	}

	// Check right cell
	if j+1 < len(board[0]) {
		colorUsingDFS(board, i, j+1)
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}
