package main

import (
	. "main/pkg/list_node"
)

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n)
// M: O(1)
// -- start --

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	count := 1

	for cur.Next != nil {
		cur = cur.Next
		count++
	}

	// circle link listed
	cur.Next = head

	k = k % count
	cur = head
	for i := 0; i < count-k-1; i++ {
		cur = cur.Next
	}

	res := cur.Next
	cur.Next = nil

	return res
}

// -- end --
