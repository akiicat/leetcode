package main
import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

func (l *ListNode) Display() {
    list := l
    for list != nil {
        fmt.Printf("%+v->", list.Val)
        list = list.Next
    }
    fmt.Printf("NULL")
    fmt.Println()
}

func main() {
  t := &ListNode{Val: 1}
  t.Next = &ListNode{Val: 2}
  t.Next.Next = &ListNode{Val: 3}
  t.Next.Next.Next = &ListNode{Val: 4}
  t.Next.Next.Next.Next = &ListNode{Val: 5}
  fmt.Printf("Input:  ")
  t.Display()
  fmt.Printf("Output: ")
  reverseList(t).Display()
  fmt.Printf("Expect: 5->4->3->2->1->NULL\n")

}

// T: O(N)
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

