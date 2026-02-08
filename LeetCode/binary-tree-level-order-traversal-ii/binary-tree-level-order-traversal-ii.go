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

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodesPerLevel := []int{}
		sizePerLevel := len(queue)

		for i := 0; i < sizePerLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			nodesPerLevel = append(nodesPerLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append([][]int{nodesPerLevel}, result...)
	}

	return result
}

func main() {
	binTree := &TreeNode{
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

	fmt.Println(levelOrderBottom(binTree))
}
