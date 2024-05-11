package main

import "fmt"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		nodesInCurrentLevel := len(queue)
		rightMostVal := 0
		for i := 0; i < nodesInCurrentLevel; i++ {
			curr := queue[0]
			queue = queue[1:]
			rightMostVal = curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, rightMostVal)
	}

	return result
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
	vertices := []int{1, 2, 3, -1, 5, 6, -1, 4}
	tree := buildBinaryTree(vertices)
	fmt.Println(rightSideView(&tree))
	return
}
