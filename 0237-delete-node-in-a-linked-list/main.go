package main
import "fmt"
import . "list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i, n, o := NewListNode([]int{4,5,1,9}), 5, NewListNode([]int{4,1,9})
  fmt.Printf("Input:  %s, %d\n", i.Sprintf(), n)
  deleteNode(i.Next)
  fmt.Printf("Output: %s\n", i.Sprintf())
  fmt.Printf("Expect: %s\n", o.Sprintf())

  i, n, o = NewListNode([]int{4,5,1,9}), 1, NewListNode([]int{4,5,9})
  fmt.Printf("Input:  %s, %d\n", i.Sprintf(), n)
  deleteNode(i.Next.Next)
  fmt.Printf("Output: %s\n", i.Sprintf())
  fmt.Printf("Expect: %s\n", o.Sprintf())
}

// T: O(1)
// M: O(1)
// -- start --

func deleteNode(node *ListNode) {
  *node = *node.Next
}

func deleteNodeAssign(node *ListNode) {
  node.Val, node.Next = node.Next.Val, node.Next.Next
}

// -- end --

