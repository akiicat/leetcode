package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestFindTilt(_t *testing.T) {
  i, o := NewTreeNode("1,2,3"), 1
  T(_t, S(i), S(findTilt(i)), S(o))

  i, o = NewTreeNode("1,2,3,4,5"), 9
  T(_t, S(i), S(findTilt(i)), S(o))
}

