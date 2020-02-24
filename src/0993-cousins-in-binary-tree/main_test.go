package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsSymmetric(_t *testing.T) {
  i, x, y, o := NewTreeNode("1,2,3,null,4,null,5"), 5, 4, true
  T(_t, S(i, x, y), S(isCousins(i, x, y)), S(o))

  i, x, y, o = NewTreeNode("1,2,3,null,4,null,5"), 5, 4, true
  T(_t, S(i, x, y), S(isCousins(i, x, y)), S(o))

  i, x, y, o = NewTreeNode("1,2,3,4"), 2, 3, false
  T(_t, S(i, x, y), S(isCousins(i, x, y)), S(o))

  i, x, y, o = NewTreeNode("1,2,3,null,4"), 2, 3, false
  T(_t, S(i, x, y), S(isCousins(i, x, y)), S(o))
}

