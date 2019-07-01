package main
import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

func main() {
  t1 := &ListNode{Val: 1}
  t1.Next = &ListNode{Val: 2}
  fmt.Printf("Input:  1->2\nOutput: %t\nExpect: false\n", isPalindrome(t1))

  t2 := &ListNode{Val: 1}
  t2.Next = &ListNode{Val: 2}
  t2.Next.Next = &ListNode{Val: 2}
  t2.Next.Next.Next = &ListNode{Val: 1}
  fmt.Printf("Input:  1->2->2->1\nOutput: %t\nExpect: true\n", isPalindrome(t2))

  t3 := &ListNode{Val: 1}
  t3.Next = &ListNode{Val: 2}
  t3.Next.Next = &ListNode{Val: 1}
  fmt.Printf("Input:  1->2->1\nOutput: %t\nExpect: true\n", isPalindrome(t3))

  t4 := &ListNode{Val: 1}
  t4.Next = &ListNode{Val: 0}
  t4.Next.Next = &ListNode{Val: 0}
  fmt.Printf("Input:  1->0->0\nOutput: %t\nExpect: false\n", isPalindrome(t4))
}

// T: O(N)
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

