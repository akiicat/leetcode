package main

import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestCountCharacters(t *testing.T) {
  i, o := NewListNode([]int{1,2,3,4,5}), NewListNode([]int{3,4,5})
  T(t, S(i), S(middleNode(i)), S(o))

  i, o = NewListNode([]int{1,2,3,4,5,6}), NewListNode([]int{4,5,6})
  T(t, S(i), S(middleNode(i)), S(o))
}

