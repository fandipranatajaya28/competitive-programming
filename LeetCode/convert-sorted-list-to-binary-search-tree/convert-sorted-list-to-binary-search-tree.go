package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}

	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	root := &TreeNode{
		Val: slow.Val,
	}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}

func inOrderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	inOrderPrint(node.Left)
	fmt.Println(node.Val)
	inOrderPrint(node.Right)
}

func constructLinkedList(list []int) *ListNode {
	linkedList := &ListNode{}
	tail := linkedList
	for _, elem := range list {
		tail.Next = &ListNode{
			Val: elem,
		}
		tail = tail.Next
	}
	return linkedList.Next
}

func main() {
	obj := constructLinkedList([]int{-10, -3, 0, 5, 9})
	inOrderPrint(sortedListToBST(obj))
}
