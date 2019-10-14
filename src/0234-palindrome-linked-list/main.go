package main
import "fmt"
import . "list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i, o := NewListNode([]int{1,2}), false
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = NewListNode([]int{1,2,2,1}), true
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = NewListNode([]int{1,2,1}), true
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = NewListNode([]int{1,0,0}), false
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)
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

