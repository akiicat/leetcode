package main
import . "main/pkg/list_node"

// T: O(n)
// M: O(1)
// -- start --

// Two Pointer
// T: O(n)
// M: O(1)
func hasCycle(head *ListNode) bool {
  if head == nil {
    return false
  }

  slow, fast := head, head
  for {
    fast = fast.Next
    if fast == nil {
      return false
    }

    fast = fast.Next
    if fast == nil {
      return false
    }

    slow = slow.Next
    if slow == fast {
      return true
    }
  }

  return false
}

// Hash
// T: O(n)
// M: O(n)
func hasCycleHash(head *ListNode) bool {
  m := make(map[*ListNode]bool)

  for head != nil {
    _, ok := m[head]
    if ok {
      return true
    }

    m[head] = true

    head = head.Next
  }

  return false
}

// -- end --

