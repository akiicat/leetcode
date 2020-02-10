package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(1) for 30
// M: O(1)
// -- start --

func getDecimalValue(head *ListNode) int {
  m := 0

  for head != nil {
    m = m << 1 | head.Val
    head = head.Next
  }

  return m
}

// -- end --

