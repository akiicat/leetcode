package main
import "fmt"
import . "list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i, o := NewListNode([]int{1,2,3,4,5}), NewListNode([]int{5,4,3,2,1})
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", reverseList(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(1)
// -- start --

func reverseList(head *ListNode) *ListNode {
  var tail *ListNode
  for head != nil {
    tail, head, head.Next = head, head.Next, tail
  }
  return tail
}

// -- end --

