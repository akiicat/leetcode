package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestAddTwoNumbers(t *testing.T) {
  i1, i2, o := NewListNode([]int{7,2,4,3}), NewListNode([]int{5,6,4}) , NewListNode([]int{7,8,0,7})
  T(t, S(i1, i2), S(addTwoNumbers(i1, i2)), S(o))
}

