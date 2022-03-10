package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestFindSecondMinimumValue(_t *testing.T) {
  i, o := NewTreeNode("2,2,5,null,null,5,7"), 5
  T(_t, S(i), S(findSecondMinimumValue(i)), S(o))

  i, o = NewTreeNode("2,2,2"), -1
  T(_t, S(i), S(findSecondMinimumValue(i)), S(o))

  i, o = NewTreeNode("5,8,5"), 8
  T(_t, S(i), S(findSecondMinimumValue(i)), S(o))
}

