package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestTrimBST(_t *testing.T) {
  i, l, r, o := NewTreeNode("1,0,2"), 1, 2, NewTreeNode("1,null,2")
  T(_t, S(i, l, r), S(trimBST(i, l, r)), S(o))

  i, l, r, o = NewTreeNode("3,0,4,null,2,null,null,1"), 1, 3, NewTreeNode("3,2,null,1")
  T(_t, S(i, l, r), S(trimBST(i, l, r)), S(o))
}

