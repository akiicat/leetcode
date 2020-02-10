package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n)
// M: O(1)
// -- start --

func middleNode(head *ListNode) *ListNode {
  count := 0
  mid := head

  for head != nil {
    if count & 0x1 == 1 {
      mid = mid.Next
    }

    count++
    head = head.Next
  }

  return mid
}

// -- end --

