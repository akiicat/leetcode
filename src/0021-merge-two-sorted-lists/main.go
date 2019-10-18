package main
import "fmt"
import . "main/pkg/list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i1, i2, o := NewListNode([]int{1,2,4}), NewListNode([]int{1,3,4}), NewListNode([]int{1,1,2,3,4,4})
  fmt.Printf("Input:  %s, %s\n", i1.ToStr(), i2.ToStr())
  fmt.Printf("Output: %s\n", mergeTwoLists(i1, i2).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(N) total nodes
// M: O(1)
// -- start --

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  head := &ListNode{}
  cur := head

  for l1 != nil && l2 != nil {

    if l1.Val < l2.Val {
      cur.Next = l1
      l1 = l1.Next
    } else {
      cur.Next = l2
      l2 = l2.Next
    }

    cur = cur.Next
  }

  if l1 == nil {
    cur.Next = l2
  } else if l2 == nil {
    cur.Next = l1
  }

  return head.Next
}

func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

// -- end --

