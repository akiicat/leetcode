package list_node
import "fmt"
import "strings"

type ListNode struct {
  Val int
  Next *ListNode
}

func NewListNode(nums []int) *ListNode {
  if len(nums) == 0 {
    return nil
  }

  head := make([]ListNode, len(nums))

  for i, num := range nums {
    head[i].Val = num
    if i < len(nums) - 1 {
      head[i].Next = &head[i + 1]
    }
  }

  return &head[0]
}

func (l *ListNode) Printf() {
  list := l
  for list != nil {
    fmt.Printf("%+v->", list.Val)
    list = list.Next
  }
  fmt.Printf("NULL")
}

func (l *ListNode) Sprintf() string {
  var s strings.Builder
  list := l
  for list != nil {
    fmt.Fprintf(&s, "%+v->", list.Val)
    list = list.Next
  }
  fmt.Fprintf(&s, "NULL")
  return s.String()
}

