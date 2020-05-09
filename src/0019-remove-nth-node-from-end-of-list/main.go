package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n)
// M: O(1)
// -- start --

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  cur, peek := head, head
  for i := 0; i < n; i++ {
    peek = peek.Next
  }

  if peek == nil {
    return head.Next
  }

  for peek.Next != nil {
    cur, peek = cur.Next, peek.Next
  }

  cur.Next = cur.Next.Next

  return head
}

// -- end --

