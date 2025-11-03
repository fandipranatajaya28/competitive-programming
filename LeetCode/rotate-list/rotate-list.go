package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Find length of the list
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	// Unnecessary work if the rotate result the same with the first
	if k%n == 0 {
		return head
	}

	modRotation := k % n

	// Make a ring
	tail.Next = head

	// New tail at (n-k-1) from head; new head = newTail.Next
	steps := n - modRotation - 1
	newTail := head
	for i := 0; i < steps; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
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
	listNum := []int{1, 2, 3, 4, 5}
	list := construct(listNum)
	result := rotateRight(list, 2)
	printList(result)
}
