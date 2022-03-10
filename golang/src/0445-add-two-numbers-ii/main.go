package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n + m)
// M: O(n + m)
// -- start --

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

  s1 := []int{}
  s2 := []int{}

  for l1 != nil {
    s1 = append(s1, l1.Val)
    l1 = l1.Next
  }
  for l2 != nil {
    s2 = append(s2, l2.Val)
    l2 = l2.Next
  }

  var head *ListNode
  c := 0
  for len(s1) != 0 || len(s2) != 0 {

    v1, v2 := 0, 0

    if len(s1) != 0 {
      v1, s1 = s1[len(s1)-1], s1[:len(s1)-1]
    }
    if len(s2) != 0 {
      v2, s2 = s2[len(s2)-1], s2[:len(s2)-1]
    }

    c = c + v1 + v2

    head = &ListNode{c%10,head}
    c /= 10
  }

  if c != 0 {
    head = &ListNode{c,head}
  }

  return head
}

// -- end --

