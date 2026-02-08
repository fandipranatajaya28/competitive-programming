package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left)
	if left == -1 {
		return -1
	}

	right := dfs(node.Right)
	if right == -1 {
		return -1
	}

	diff := abs(left - right)
	if diff > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(isBalanced(root))
}
