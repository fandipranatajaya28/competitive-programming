package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var (
		levelOrderNodes [][]int
		queue           []*TreeNode
	)

	if root == nil {
		return levelOrderNodes
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		nodesPerLevel := []int{}
		sizePerLevel := len(queue)

		for range sizePerLevel {
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

		levelOrderNodes = append(levelOrderNodes, nodesPerLevel)
	}

	return levelOrderNodes
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

	fmt.Println(levelOrder(binTree))
}
