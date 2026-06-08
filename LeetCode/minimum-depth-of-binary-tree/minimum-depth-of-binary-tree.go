package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}

// 	if root.Left == nil {
// 		return 1 + minDepth(root.Right)
// 	}

// 	if root.Right == nil {
// 		return 1 + minDepth(root.Left)
// 	}

// 	return 1 + min(minDepth(root.Left), minDepth(root.Right))
// }

func minDepth(root *TreeNode) int {
	// Empty tree has depth 0.
	if root == nil {
		return 0
	}

	// Queue stores nodes to process.
	queue := []*TreeNode{root}

	// Root is depth 1.
	depth := 1

	for len(queue) > 0 {
		// Process one level at a time.
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// First leaf found by BFS gives the minimum depth.
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// Add left child if exists.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// Add right child if exists.
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Move to next level.
		depth++
	}

	return depth
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
	fmt.Println(minDepth(root))
}
