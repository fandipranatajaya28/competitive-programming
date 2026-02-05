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

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 9,
				},
			},
		},
	}
	fmt.Println(inorderTraversal(tree))
	// Output: [4,2,6,5,7,1,3,9,8]
}
