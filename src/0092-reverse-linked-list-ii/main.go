package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n)
// M: O(1)
// -- start --

func reverseBetween(head *ListNode, m int, n int) *ListNode {
  var prev, cur, next *ListNode
  prev, cur, next = nil, &ListNode{0,head}, head
  head = cur

  for i := 0; i < m-1; i++ {
    prev, cur, next = cur, next, next.Next
  }

  l := cur
  for i := 0; i < n-m+1; i++ {
    prev, cur, next = cur, next, next.Next
    cur.Next = prev
  }

  l.Next.Next = next
  l.Next = cur

  return head.Next
}

// -- end --

