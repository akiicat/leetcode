package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestMergeTwoLists(_t *testing.T) {
  i1, i2, o := NewListNode([]int{1,2,4}), NewListNode([]int{1,3,4}), NewListNode([]int{1,1,2,3,4,4})
  T(_t, S(i1, i2), S(mergeTwoLists(i1, i2)), S(o))

  i1, i2, o = NewListNode([]int{1,3,4}), NewListNode([]int{1,2,4,5}), NewListNode([]int{1,1,2,3,4,4,5})
  T(_t, S(i1, i2), S(mergeTwoLists(i1, i2)), S(o))

  i1, i2, o = NewListNode([]int{1,3,4,5}), NewListNode([]int{1,2,4}), NewListNode([]int{1,1,2,3,4,4,5})
  T(_t, S(i1, i2), S(mergeTwoLists(i1, i2)), S(o))
}

func TestMergeTwoListsRecursive(_t *testing.T) {
  i1, i2, o := NewListNode([]int{1,2,4}), NewListNode([]int{1,3,4}), NewListNode([]int{1,1,2,3,4,4})
  T(_t, S(i1, i2), S(mergeTwoListsRecursive(i1, i2)), S(o))

  i1, i2, o = NewListNode([]int{1,3,4}), NewListNode([]int{1,2,4,5}), NewListNode([]int{1,1,2,3,4,4,5})
  T(_t, S(i1, i2), S(mergeTwoListsRecursive(i1, i2)), S(o))

  i1, i2, o = NewListNode([]int{1,3,4,5}), NewListNode([]int{1,2,4}), NewListNode([]int{1,1,2,3,4,4,5})
  T(_t, S(i1, i2), S(mergeTwoListsRecursive(i1, i2)), S(o))
}

