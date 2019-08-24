package main
import "fmt"
import . "list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i, o := NewListNode([]int{1,1,2}), NewListNode([]int{1,2})
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", deleteDuplicates(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i, o = NewListNode([]int{1,1,2,3,3}), NewListNode([]int{1,2,3})
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", deleteDuplicates(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i, o = NewListNode([]int{1,1,1}), NewListNode([]int{1})
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", deleteDuplicates(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i, o = NewListNode([]int{}), NewListNode([]int{})
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", deleteDuplicates(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

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
