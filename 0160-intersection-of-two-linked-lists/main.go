package main
import "fmt"
import . "list_node"

// type ListNode struct {
//   Val int
//   Next *ListNode
// }

func main() {
  i1, i2, n1, n2, n := NewListNode([]int{4,1,8,4,5}), NewListNode([]int{5,0,1,8,4,5}), 2, 3, 8
  o := intersect(i1, i2, n1, n2)
  fmt.Printf("Input:  %s, %s, %d, %d, %d\n", i1.ToStr(), i2.ToStr(), n1, n2, n)
  fmt.Printf("Output: %s\n", getIntersectionNode(i1, i2).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i1, i2, n1, n2, n = NewListNode([]int{0,9,1,2,4}), NewListNode([]int{3,2,4}), 3, 1, 2
  o = intersect(i1, i2, n1, n2)
  fmt.Printf("Input:  %s, %s, %d, %d, %d\n", i1.ToStr(), i2.ToStr(), n1, n2, n)
  fmt.Printf("Output: %s\n", getIntersectionNode(i1, i2).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i1, i2, n1, n2, n = NewListNode([]int{2,6,4}), NewListNode([]int{1,5}), 3, 2, 0
  o = intersect(i1, i2, n1, n2)
  fmt.Printf("Input:  %s, %s, %d, %d, %d\n", i1.ToStr(), i2.ToStr(), n1, n2, n)
  fmt.Printf("Output: %s\n", getIntersectionNode(i1, i2).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i1, i2, n1, n2, n = NewListNode([]int{2,6,4}), NewListNode([]int{1,5,2,2}), 3, 4, 0
  o = intersect(i1, i2, n1, n2)
  fmt.Printf("Input:  %s, %s, %d, %d, %d\n", i1.ToStr(), i2.ToStr(), n1, n2, n)
  fmt.Printf("Output: %s\n", getIntersectionNode(i1, i2).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

func intersect(i1, i2 *ListNode, n1, n2 int) *ListNode {
  r := i1
  for i := 0; i < n1 - 1; i++ {
    r = r.Next
  }

  l := i2
  for i := 0; i < n2 - 1; i++ {
    l = l.Next
  }

  l.Next = r.Next

  return l.Next
}

// T: O(n + m)
// M: O(1)
// -- start --

// T: O(n + m)
// M: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {

  // listA = A + Common
  // listB = B + Common
  // PointerA: [A + Common + B] + Common
  // PointerB: [B + Common + A] + Common
  if headA == nil || headB == nil {
    return nil
  }
  a := headA
  b := headB
  for  a != b {
    if a == nil {
      a = headB
    }else{
      a = a.Next
    }
    if b == nil {
      b = headA
    }else{
      b = b.Next
    }
  }
  return a
}

// T: O(n + m)
// M: O(n)
func getIntersectionNodeMap(headA, headB *ListNode) *ListNode {
  m := make(map[*ListNode]bool)

  for headA != nil {
    m[headA] = true
    headA = headA.Next
  }

  for headB != nil {
    _, ok := m[headB]
    if ok {
      return headB
    }
    headB = headB.Next
  }

  return nil
}

// -- end --

