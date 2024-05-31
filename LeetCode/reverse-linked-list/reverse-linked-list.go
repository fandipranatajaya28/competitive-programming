package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func generateList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	// Initialize the head of the linked list
	head := &ListNode{Val: arr[0]}
	current := head

	// Iterate over the array and create the linked list
	for _, value := range arr[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return head
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := generateList(arr)
	ans := reverseList(list)
	fmt.Println(ans)
	return
}
