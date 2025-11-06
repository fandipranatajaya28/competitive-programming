package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := divideAndGetMidList(head)
	left := sortList(head)
	right := sortList(mid)

	return mergeList(left, right)
}

func divideAndGetMidList(head *ListNode) *ListNode {
	var prev *ListNode

	// slow move 1 step, fast move 2 step
	// when fast reach the end of list, slow will be in the middle

	/*
		Imagine two people — Slow and Fast — running on the same straight path (your linked list).
		- Slow runs 1 step per second.
		- Fast runs 2 steps per second.
		They both start from the same point.

		Now, if Fast reaches the end of the path (the tail of the list),
		then Slow must have covered half the distance,
		because Fast moves twice as fast.

		So when Fast has finished the race,
		→ Slow is at the middle of the track.
	*/

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// To divide the List
	prev.Next = nil

	return slow
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
	mergedList := &ListNode{}
	tail := mergedList

	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}
		tail = tail.Next
	}

	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}

	return mergedList.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	obj := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	printList(sortList(obj))
}
