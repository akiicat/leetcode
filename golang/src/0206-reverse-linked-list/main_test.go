package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestReverseList(_t *testing.T) {
  i, o := NewListNode([]int{1,2,3,4,5}), NewListNode([]int{5,4,3,2,1})
  T(_t, S(i), S(reverseList(i)), S(o))
}
