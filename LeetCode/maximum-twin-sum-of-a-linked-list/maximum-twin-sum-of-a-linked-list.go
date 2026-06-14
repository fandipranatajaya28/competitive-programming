package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Step 1:
	// Use slow and fast pointers to find the middle of the linked list.
	//
	// Since the list length is even:
	// - slow will stop at the first node of the second half.
	//
	// Example:
	// 1 -> 2 -> 3 -> 4
	//           ^
	//          slow
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2:
	// Reverse the second half of the linked list.
	//
	// Example:
	// first half:  1 -> 2
	// second half: 3 -> 4
	//
	// reversed second half: 4 -> 3
	secondHalf := reverseList(slow)

	// Step 3:
	// Compare first half and reversed second half.
	//
	// Example:
	// first half:           1 -> 2
	// reversed second half: 4 -> 3
	//
	// twin sums:
	// 1 + 4 = 5
	// 2 + 3 = 5
	maxSum := 0

	first := head
	second := secondHalf

	for second != nil {
		sum := first.Val + second.Val

		if sum > maxSum {
			maxSum = sum
		}

		first = first.Next
		second = second.Next
	}

	return maxSum
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		// Save next node before changing curr.Next.
		nextNode := curr.Next

		// Reverse pointer.
		curr.Next = prev

		// Move prev and curr forward.
		prev = curr
		curr = nextNode
	}

	// prev is the new head of the reversed list.
	return prev
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

	result := pairSum(head)

	fmt.Println("Maximum twin sum:", result)
}
