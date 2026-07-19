package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Symmetric Tree:
// Return true if a binary tree is a mirror image of itself around its center.
//
// Key insight: the tree is symmetric iff its left and right subtrees are
// MIRRORS of each other. Two trees mirror when:
//   1. both are nil, OR
//   2. both are non-nil, values are equal, AND — the crossed part —
//      left-of-one mirrors right-of-the-other, and right-of-one mirrors
//      left-of-the-other.
//
// Time: O(n). Space: O(h) recursion stack (h = height).
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror reports whether trees a and b are mirror images of each other.
func isMirror(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	// Cross-compare: a's left with b's right, a's right with b's left.
	return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}

func main() {
	// Tree 1 (symmetric):
	//        1
	//       / \
	//      2   2
	//     / \ / \
	//    3  4 4  3
	symmetric := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}

	// Tree 2 (NOT symmetric):
	//        1
	//       / \
	//      2   2
	//       \   \
	//        3   3
	notSymmetric := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}

	fmt.Println(isSymmetric(symmetric))    // true
	fmt.Println(isSymmetric(notSymmetric)) // false
}
