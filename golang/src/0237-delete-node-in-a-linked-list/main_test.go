package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestDeleteNode(_t *testing.T) {
  i, n, o := NewListNode([]int{4,5,1,9}), 5, NewListNode([]int{4,1,9})
  _i := S(i, n)
  deleteNode(i.Next)
  T(_t, _i, S(i), S(o))

  i, n, o = NewListNode([]int{4,5,1,9}), 1, NewListNode([]int{4,5,9})
  _i = S(i, n)
  deleteNode(i.Next.Next)
  T(_t, _i, S(i), S(o))
}

func TestDeleteNodeAssign(_t *testing.T) {
  i, n, o := NewListNode([]int{4,5,1,9}), 5, NewListNode([]int{4,1,9})
  _i := S(i, n)
  deleteNodeAssign(i.Next)
  T(_t, _i, S(i), S(o))

  i, n, o = NewListNode([]int{4,5,1,9}), 1, NewListNode([]int{4,5,9})
  _i = S(i, n)
  deleteNodeAssign(i.Next.Next)
  T(_t, _i, S(i), S(o))
}
