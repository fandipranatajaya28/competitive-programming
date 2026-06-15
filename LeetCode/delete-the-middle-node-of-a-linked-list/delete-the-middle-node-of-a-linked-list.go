package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	// If only one node, deleting it returns empty list.
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	var prev *ListNode

	// Move fast 2 steps and slow 1 step.
	// When fast reaches the end, slow is at the middle node.
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Delete middle node.
	prev.Next = slow.Next

	return head
}

// Helper function to create a linked list from an array.
func buildLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head

	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}

	return head
}

// Helper function to print linked list.
// Not needed for LeetCode, only for local testing.
func printLinkedList(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Print(curr.Val)

		if curr.Next != nil {
			fmt.Print(" -> ")
		}

		curr = curr.Next
	}

	fmt.Println()
}

func main() {
	nums := []int{5, 4, 2, 1}

	head := buildLinkedList(nums)

	fmt.Print("Linked list: ")
	printLinkedList(head)

	result := deleteMiddle(head)

	fmt.Println("Result:", result)
}
