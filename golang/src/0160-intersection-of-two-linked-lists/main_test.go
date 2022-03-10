package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestGetIntersectionNode(_t *testing.T) {
  i1, i2, n1, n2, n := NewListNode([]int{4,1,8,4,5}), NewListNode([]int{5,0,1,8,4,5}), 2, 3, 8
  o := intersect(i1, i2, n1, n2)
  T(_t, S(i1, i2, n1, n2, n), S(getIntersectionNode(i1, i2)), S(o))

  i1, i2, n1, n2, n = NewListNode([]int{0,9,1,2,4}), NewListNode([]int{3,2,4}), 3, 1, 2
  o = intersect(i1, i2, n1, n2)
  T(_t, S(i1, i2, n1, n2, n), S(getIntersectionNode(i1, i2)), S(o))

  i1, i2, n1, n2, n = NewListNode([]int{2,6,4}), NewListNode([]int{1,5}), 3, 2, 0
  o = intersect(i1, i2, n1, n2)
  T(_t, S(i1, i2, n1, n2, n), S(getIntersectionNode(i1, i2)), S(o))

  i1, i2, n1, n2, n = NewListNode([]int{2,6,4}), NewListNode([]int{1,5,2,2}), 3, 4, 0
  o = intersect(i1, i2, n1, n2)
  T(_t, S(i1, i2, n1, n2, n), S(getIntersectionNode(i1, i2)), S(o))
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

