package main
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

// T: O(n)
// M: O(1)
// -- start --

func isPalindrome(head *ListNode) bool {

  if head == nil {
    return true
  }

  // go to the middle (two pointer method)    
  // reverse second half    
  // compare first half with reverse second half
  // if equal we have palindrome, otherwise not

  slow := head
  fast := head.Next

  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }

  secondHalf := reverseList(slow)

  for head != nil && secondHalf != nil {
    if head.Val != secondHalf.Val {
      return false
    }

    head = head.Next
    secondHalf = secondHalf.Next
  }

  return true
}

func reverseList(head *ListNode) *ListNode {
  var tail *ListNode

  for head != nil {
    head.Next, tail, head = tail, head, head.Next
  }

  return tail
}

// -- end --

