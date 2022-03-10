package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestDeleteDuplicates(_t *testing.T) {
  i, o := NewListNode([]int{1,1,2}), NewListNode([]int{1,2})
  T(_t, S(i), S(deleteDuplicates(i)), S(o))

  i, o = NewListNode([]int{1,1,2,3,3}), NewListNode([]int{1,2,3})
  T(_t, S(i), S(deleteDuplicates(i)), S(o))

  i, o = NewListNode([]int{1,1,1}), NewListNode([]int{1})
  T(_t, S(i), S(deleteDuplicates(i)), S(o))

  i, o = NewListNode([]int{}), NewListNode([]int{})
  T(_t, S(i), S(deleteDuplicates(i)), S(o))
}
