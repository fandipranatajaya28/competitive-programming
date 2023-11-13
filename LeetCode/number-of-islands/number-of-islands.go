package main

import "fmt"

func isValid(row int, col int, grid [][]byte, visited [][]bool) bool {
	if row >= len(grid) || row < 0 {
		return false
	} else if col >= len(grid[0]) || col < 0 {
		return false
	} else if grid[row][col] == '0' {
		return false
	} else if visited[row][col] {
		return false
	}

	return true
}

func bfs(row int, col int, grid [][]byte, visited [][]bool) {
	var (
		directions = [][]int{
			// row, col
			{1, 0},  // down
			{-1, 0}, // up
			{0, 1},  // right
			{0, -1}, // left
		}
		queue = [][]int{}
	)

	queue = append(queue, []int{row, col})
	visited[row][col] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// traverse the neighbours
		for _, direction := range directions {
			r := current[0] + direction[0]
			c := current[1] + direction[1]
			if isValid(r, c, grid, visited) {
				queue = append(queue, []int{r, c})
				visited[r][c] = true
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	// initiate visited coordinates
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	numOfIslands := 0

	if len(grid) > 0 {
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if grid[row][col] == '1' && !visited[row][col] {
					bfs(row, col, grid, visited)
					numOfIslands++
				}
			}
		}
	}

	return numOfIslands
}

func main() {
	grid := [][]byte{
		{
			'1', '1', '0', '0', '0',
		},
		{
			'1', '1', '0', '0', '0',
		},
		{
			'0', '0', '1', '0', '0',
		},
		{
			'0', '0', '0', '1', '1',
		},
	}
	fmt.Println(numIslands(grid))
	// output: 3
}
