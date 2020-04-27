package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestAddTwoNumbers(_t *testing.T) {
  l1, l2, o := NewListNode([]int{2,4,3}), NewListNode([]int{5,6,4}), NewListNode([]int{7,0,8})
  T(_t, S(l1, l2), S(addTwoNumbers(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{5,9}), NewListNode([]int{5}), NewListNode([]int{0,0,1})
  T(_t, S(l1, l2), S(addTwoNumbers(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{8,9,9}), NewListNode([]int{2}), NewListNode([]int{0,0,0,1})
  T(_t, S(l1, l2), S(addTwoNumbers(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{0}), NewListNode([]int{7,3}), NewListNode([]int{7,3})
  T(_t, S(l1, l2), S(addTwoNumbers(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{5}), NewListNode([]int{5}), NewListNode([]int{0,1})
  T(_t, S(l1, l2), S(addTwoNumbers(l1, l2)), S(o))
}

func TestAddTwoNumbersLoop(_t *testing.T) {
  l1, l2, o := NewListNode([]int{2,4,3}), NewListNode([]int{5,6,4}), NewListNode([]int{7,0,8})
  T(_t, S(l1, l2), S(addTwoNumbersLoop(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{5,9}), NewListNode([]int{5}), NewListNode([]int{0,0,1})
  T(_t, S(l1, l2), S(addTwoNumbersLoop(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{8,9,9}), NewListNode([]int{2}), NewListNode([]int{0,0,0,1})
  T(_t, S(l1, l2), S(addTwoNumbersLoop(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{0}), NewListNode([]int{7,3}), NewListNode([]int{7,3})
  T(_t, S(l1, l2), S(addTwoNumbersLoop(l1, l2)), S(o))

  l1, l2, o = NewListNode([]int{5}), NewListNode([]int{5}), NewListNode([]int{0,1})
  T(_t, S(l1, l2), S(addTwoNumbersLoop(l1, l2)), S(o))
}

