package main
import "fmt"
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i, n, o := NewListNode([]int{1,2,6,3,4,5,6}), 6, NewListNode([]int{1,2,3,4,5})
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), n)
  fmt.Printf("Output: %s\n", removeElements(i, n).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i, n, o = NewListNode([]int{1}), 1, NewListNode([]int{})
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), n)
  fmt.Printf("Output: %s\n", removeElements(i, n).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i, n, o = NewListNode([]int{1, 1}), 1, NewListNode([]int{})
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), n)
  fmt.Printf("Output: %s\n", removeElements(i, n).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(1)
// -- start --

func removeElements(head *ListNode, val int) *ListNode {
  head_prev := &ListNode{Next: head}

  for cur, prev := head, head_prev; cur != nil; cur, prev = cur.Next, cur {
    if cur.Val == val {
      prev.Next, cur = cur.Next, prev
    }
  }

  return head_prev.Next
}

// -- end --

