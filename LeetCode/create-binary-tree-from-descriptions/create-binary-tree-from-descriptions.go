package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	// nodeMap[val] stores the TreeNode pointer for value val.
	// This helps us reuse the same node whenever we see the same value again.
	nodeMap := make(map[int]*TreeNode)

	// isChild[val] = true means val appears as a child of another node.
	// The root is the only node that never appears as a child.
	isChild := make(map[int]bool)

	for _, desc := range descriptions {
		parentVal := desc[0]
		childVal := desc[1]
		isLeft := desc[2]

		// Create parent node if it does not exist yet.
		if _, exists := nodeMap[parentVal]; !exists {
			nodeMap[parentVal] = &TreeNode{Val: parentVal}
		}

		// Create child node if it does not exist yet.
		if _, exists := nodeMap[childVal]; !exists {
			nodeMap[childVal] = &TreeNode{Val: childVal}
		}

		parentNode := nodeMap[parentVal]
		childNode := nodeMap[childVal]

		// Connect child to parent's left or right.
		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}

		// Mark this value as a child.
		isChild[childVal] = true
	}

	// Find the node that never appears as a child.
	// That node is the root of the tree.
	for val, node := range nodeMap {
		if !isChild[val] {
			return node
		}
	}

	return nil
}

func main() {
	createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}})
}
