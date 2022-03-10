package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestGetMinimumDifference(_t *testing.T) {
  i, o := NewTreeNode("1,null,3,2"), 1
  T(_t, S(i), S(getMinimumDifference(i)), S(o))

  i, o = NewTreeNode("5,4,7"), 1
  T(_t, S(i), S(getMinimumDifference(i)), S(o))
}

