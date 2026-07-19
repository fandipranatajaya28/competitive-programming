package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Intersection of Two Linked Lists:
// Return the node where two singly-linked lists intersect, or nil if they don't.
// Intersection is by REFERENCE (same node), not by equal value.
//
// Two-pointer trick: start pA at headA and pB at headB. Advance both one step
// at a time. Whenever a pointer reaches the end, redirect it to the OTHER list's
// head. This equalizes the distance travelled:
//   pA walks lenA then lenB;  pB walks lenB then lenA.
// After at most lenA + lenB steps they either meet at the intersection node or
// both become nil simultaneously (no intersection).
//
// Why it works with a shared tail of length c and unique prefixes a and b:
//   pA: a + c + b   |   pB: b + c + a   → equal totals, so they align.
//
// Time: O(lenA + lenB). Space: O(1).
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	// Loop until they meet. They are guaranteed to converge: either at the
	// intersection node, or both at nil (which also satisfies pA == pB).
	for pA != pB {
		// When pA reaches the end, jump to headB; otherwise step forward.
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		// Symmetric for pB.
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}

func main() {
	// Build two lists that share a common tail (the node with Val 8):
	//   A: 4 -> 1 --\
	//               8 -> 4 -> 5
	//   B: 5 -> 6 -> 1 --/
	// Intersection is the node with Val 8.
	shared := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: shared}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: shared}}}

	if node := getIntersectionNode(headA, headB); node != nil {
		fmt.Println("Intersection at node with Val:", node.Val) // 8
	} else {
		fmt.Println("No intersection")
	}

	// Two lists with no shared node.
	x := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	y := &ListNode{Val: 1, Next: &ListNode{Val: 5}}

	if node := getIntersectionNode(x, y); node != nil {
		fmt.Println("Intersection at node with Val:", node.Val)
	} else {
		fmt.Println("No intersection") // this branch
	}
}
