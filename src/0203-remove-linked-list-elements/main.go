package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n)
// M: O(1)
// -- start --

func removeElements(head *ListNode, val int) *ListNode {
  head_prev := &ListNode{Next: head}

  for cur, prev := head, head_prev; cur != nil; cur, prev = cur.Next, cur {
    if cur.Val == val {
      prev.Next, cur = cur.Next, prev
    }
  }

  return head_prev.Next
}

// -- end --

