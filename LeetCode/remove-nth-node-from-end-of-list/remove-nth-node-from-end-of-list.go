package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node before the head node
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// Move fast n+1 step forward
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move slow and fast at the same time
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Delete the node pointed to by slow.Next
	slow.Next = slow.Next.Next

	// Return the updated list, skipping the dummy node
	return dummy.Next
}

func construct(values []int) *ListNode {
	var head ListNode
	curr := &head
	for idx, value := range values {
		curr.Val = value
		if idx == len(values)-1 {
			break
		}
		curr.Next = &ListNode{}
		curr = curr.Next
	}
	return &head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	head := construct([]int{1, 2, 3, 4, 5})
	printList(removeNthFromEnd(head, 2))
}
