package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestRemoveNthFromEnd(t *testing.T) {
  i, n, o := NewListNode([]int{1,2,3,4,5}), 2, NewListNode([]int{1,2,3,5})
  T(t, S(i, n), S(removeNthFromEnd(i, n)), S(o))

  i, n, o = NewListNode([]int{1}), 1, NewListNode([]int{})
  T(t, S(i, n), S(removeNthFromEnd(i, n)), S(o))
}

