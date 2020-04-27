package main
import . "main/pkg/list_node"

// T: O(n)
// M: O(1)
// -- start --

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  return add(l1, l2, 0)
}

func add(l1 *ListNode, l2 *ListNode, c int) *ListNode {
  if l1 == nil && l2 == nil {
    if c == 0 {
      return nil
    }
    return &ListNode{c, nil}
  }

  sum := c
  if l1 == nil {
    l1, l2 = l2, l1
  }
  if l2 == nil {
    sum += l1.Val
    l1.Val = sum % 10
    l1.Next = add(l1.Next, nil, sum/10)
    return l1
  }

  sum += l1.Val + l2.Val
  l1.Val = sum % 10
  l1.Next = add(l1.Next, l2.Next, sum/10)
  return l1
}

func addTwoNumbersLoop(l1 *ListNode, l2 *ListNode) *ListNode {

  head := l1

  for l1 != nil && l2 != nil {
    l1.Val += l2.Val
    l2.Val = 0

    if l1.Next == nil && l2.Next != nil {
      l1.Next, l2.Next = l2.Next, l1.Next
    }

    if l1.Val >= 10 {
      if l1.Next == nil {
        l1.Next = &ListNode{}
      }

      l1.Next.Val += l1.Val / 10
      l1.Val %= 10
    }

    l1, l2 = l1.Next, l2.Next
  }

  for l1 != nil && l1.Val >= 10 {
    if l1.Next == nil {
      l1.Next = &ListNode{}
    }

    l1.Next.Val += l1.Val / 10
    l1.Val %= 10

    l1 = l1.Next
  }

  return head
}

// -- end --

