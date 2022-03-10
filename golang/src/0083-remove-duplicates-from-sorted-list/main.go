package main
import . "main/pkg/list_node"

// leetcode 26.
// T: O(n)
// M: O(1)
// -- start --

func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }
  cur := head

  for cur.Next != nil {
    if cur.Val ==  cur.Next.Val {
      cur.Next = cur.Next.Next
    } else {
      cur = cur.Next
    }
  }

  return head
}

// -- end --
