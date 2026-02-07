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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	mergedTail := mergedList
	tail1 := list1
	tail2 := list2

	for tail1 != nil && tail2 != nil {
		if tail1.Val <= tail2.Val {
			mergedTail.Next = &ListNode{
				Val: tail1.Val,
			}
			mergedTail = mergedTail.Next
			tail1 = tail1.Next
		} else {
			mergedTail.Next = &ListNode{
				Val: tail2.Val,
			}
			mergedTail = mergedTail.Next
			tail2 = tail2.Next
		}
	}

	for tail1 != nil {
		mergedTail.Next = &ListNode{
			Val: tail1.Val,
		}
		mergedTail = mergedTail.Next
		tail1 = tail1.Next
	}

	for tail2 != nil {
		mergedTail.Next = &ListNode{
			Val: tail2.Val,
		}
		mergedTail = mergedTail.Next
		tail2 = tail2.Next
	}

	return mergedList.Next
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
	printList(mergeTwoLists(construct([]int{1, 2, 4}), construct([]int{1, 3, 4})))
}
