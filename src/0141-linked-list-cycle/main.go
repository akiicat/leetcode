package main
import "fmt"
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i, pos, o := NewListNode([]int{3,2,0,-4}), 1, true
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), pos)
  AddPos(i, pos)
  fmt.Printf("Output: %t\n", hasCycle(i))
  fmt.Printf("Expect: %t\n", o)

  i, pos, o = NewListNode([]int{1,2}), 0, true
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), pos)
  AddPos(i, pos)
  fmt.Printf("Output: %t\n", hasCycle(i))
  fmt.Printf("Expect: %t\n", o)

  i, pos, o = NewListNode([]int{1,2}), -1, false
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), pos)
  AddPos(i, pos)
  fmt.Printf("Output: %t\n", hasCycle(i))
  fmt.Printf("Expect: %t\n", o)
}

func AddPos(head *ListNode, pos int) {
  if pos < 0 {
    return
  }

  cur := head

  for {
    if pos == 0 {
      cur = head
    }

    if head.Next == nil {
      head.Next = cur
      return
    }

    head = head.Next
  }
}

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

