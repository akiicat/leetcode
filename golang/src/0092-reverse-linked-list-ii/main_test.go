package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestReverseBetween(t *testing.T) {
  i, m, n, o := NewListNode([]int{1,2,3,4,5}), 2, 4, NewListNode([]int{1,4,3,2,5})
  T(t, S(i, m, n), S(reverseBetween(i, m, n)), S(o))

  i, m, n, o = NewListNode([]int{3,5}), 2, 2, NewListNode([]int{3,5})
  T(t, S(i, m, n), S(reverseBetween(i, m, n)), S(o))

  i, m, n, o = NewListNode([]int{1,2,3,4}), 1, 4, NewListNode([]int{4,3,2,1})
  T(t, S(i, m, n), S(reverseBetween(i, m, n)), S(o))
}

