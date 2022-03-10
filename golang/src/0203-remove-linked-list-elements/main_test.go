package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestRemoveElements(_t *testing.T) {
  i, n, o := NewListNode([]int{1,2,6,3,4,5,6}), 6, NewListNode([]int{1,2,3,4,5})
  T(_t, S(i, n), S(removeElements(i, n)), S(o))

  i, n, o = NewListNode([]int{1}), 1, NewListNode([]int{})
  T(_t, S(i, n), S(removeElements(i, n)), S(o))

  i, n, o = NewListNode([]int{1, 1}), 1, NewListNode([]int{})
  T(_t, S(i, n), S(removeElements(i, n)), S(o))
}

