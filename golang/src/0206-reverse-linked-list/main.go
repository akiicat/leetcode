package main
import . "main/pkg/list_node"

// T: O(n)
// M: O(1)
// -- start --

func reverseList(head *ListNode) *ListNode {
  var tail *ListNode
  for head != nil {
    tail, head, head.Next = head, head.Next, tail
  }
  return tail
}

// -- end --

