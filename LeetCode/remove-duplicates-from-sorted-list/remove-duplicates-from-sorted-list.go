package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tail := head
	for tail != nil && tail.Next != nil {
		if tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
			continue
		}
		tail = tail.Next
	}
	return head
}

func construct(values []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, val := range values {
		curr.Next = &ListNode{
			Val: val,
		}
		curr = curr.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	printList(deleteDuplicates(construct([]int{1, 1, 2, 3, 3})))
}
