package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	evenList := &ListNode{}
	evenTail := evenList
	tail := head
	for tail.Next != nil {
		evenTail.Next = &ListNode{
			Val: tail.Next.Val,
		}
		evenTail = evenTail.Next
		tail.Next = tail.Next.Next
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}
	tail.Next = evenList.Next

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func constructList(list []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, elem := range list {
		tail.Next = &ListNode{
			Val: elem,
		}
		tail = tail.Next
	}
	return head.Next
}

func main() {
	head := constructList([]int{1, 2, 3, 4, 5})
	res := oddEvenList(head)
	printList(res)
}
