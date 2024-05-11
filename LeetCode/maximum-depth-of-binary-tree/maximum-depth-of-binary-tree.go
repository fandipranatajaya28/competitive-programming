package main

import "fmt"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	max := leftDepth
	if rightDepth > max {
		max = rightDepth
	}

	return 1 + max
}

func buildBinaryTree(vertices []int) (tree TreeNode) {
	if len(vertices) == 0 {
		return
	}

	tree.Val = vertices[0]
	queue := []*TreeNode{}
	queue = append(queue, &tree)
	i := 1

	for i < len(vertices) {
		curr := queue[0]
		queue = queue[1:]
		if i < len(vertices) {
			if vertices[i] != -1 {
				curr.Left = &TreeNode{
					Val: vertices[i],
				}
				queue = append(queue, curr.Left)
			}
			i++
		}
		if i < len(vertices) {
			if vertices[i] != -1 {
				curr.Right = &TreeNode{
					Val: vertices[i],
				}
				queue = append(queue, curr.Right)
			}
			i++
		}
	}
	return
}

func main() {
	vertices := []int{3, 9, 20, -1, -1, 15, 7}
	tree := buildBinaryTree(vertices)
	fmt.Println(maxDepth(&tree))
	return
}
