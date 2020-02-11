package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n + m)
// M: O(1)
// -- start --

// T: O(n + m)
// M: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {

  // listA = A + Common
  // listB = B + Common
  // PointerA: [A + Common + B] + Common
  // PointerB: [B + Common + A] + Common
  if headA == nil || headB == nil {
    return nil
  }
  a := headA
  b := headB
  for  a != b {
    if a == nil {
      a = headB
    }else{
      a = a.Next
    }
    if b == nil {
      b = headA
    }else{
      b = b.Next
    }
  }
  return a
}

// T: O(n + m)
// M: O(n)
func getIntersectionNodeMap(headA, headB *ListNode) *ListNode {
  m := make(map[*ListNode]bool)

  for headA != nil {
    m[headA] = true
    headA = headA.Next
  }

  for headB != nil {
    _, ok := m[headB]
    if ok {
      return headB
    }
    headB = headB.Next
  }

  return nil
}

// -- end --

