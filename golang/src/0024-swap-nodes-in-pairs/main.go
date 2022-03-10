package main

import (
	. "main/pkg/list_node"
)

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n) total nodes
// M: O(1)
// -- start --

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := head.Next

	for head != nil && head.Next != nil {
		next := head.Next.Next
		if next != nil && next.Next != nil {
			next = next.Next
		}

		head, head.Next, head.Next.Next = head.Next.Next, next, head
	}

	return res
}

// -- end --
