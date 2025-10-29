package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

func preOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)

	preOrderPrint(root.Left)

	preOrderPrint(root.Right)
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
	res := searchBST(node, 2)
	preOrderPrint(res)
}
