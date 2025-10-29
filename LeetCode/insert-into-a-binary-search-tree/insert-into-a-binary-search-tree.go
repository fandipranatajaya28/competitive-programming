package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func inOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}

	inOrderPrint(root.Left)

	fmt.Println(root.Val)

	inOrderPrint(root.Right)
}

func main() {
	var node = new(TreeNode)
	node.Val = 4
	node.Left = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	node.Right = &TreeNode{
		Val: 7,
	}
	res := insertIntoBST(node, 5)
	inOrderPrint(res)
}
