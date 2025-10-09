package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
	}

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
	firstList := construct([]int{2, 4, 3})
	secondList := construct([]int{5, 6, 4})
	res := addTwoNumbers(firstList, secondList)
	printList(res)
}
