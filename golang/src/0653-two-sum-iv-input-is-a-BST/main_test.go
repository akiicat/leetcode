package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestFindTarget(_t *testing.T) {
  i, t, o := NewTreeNode("5,3,6,2,4,0,7"), 9, true
  T(_t, S(i, t), S(findTarget(i, t)), S(o))
}

