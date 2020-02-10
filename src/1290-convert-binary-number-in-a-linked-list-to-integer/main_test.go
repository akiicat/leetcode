package main

import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestGetDecimalValue(t *testing.T) {
  i, o := NewListNode([]int{1,0,1}), 5
  T(t, S(i), S(getDecimalValue(i)), S(o))

  i, o = NewListNode([]int{0}), 0
  T(t, S(i), S(getDecimalValue(i)), S(o))

  i, o = NewListNode([]int{1}), 1
  T(t, S(i), S(getDecimalValue(i)), S(o))

  i, o = NewListNode([]int{1,0,0,1,0,0,1,1,1,0,0,0,0,0,0}), 18880
  T(t, S(i), S(getDecimalValue(i)), S(o))

  i, o = NewListNode([]int{0,0}), 0
  T(t, S(i), S(getDecimalValue(i)), S(o))
}

