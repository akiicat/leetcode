package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestHasPathSum(_t *testing.T) {
  i1, i2, o := NewTreeNode("5,4,8,11,null,13,4,7,2,null,null,null,1"), 22, true
  T(_t, S(i1, i2), S(hasPathSum(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("1,2"), 0, false
  T(_t, S(i1, i2), S(hasPathSum(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("1,2"), 10, false
  T(_t, S(i1, i2), S(hasPathSum(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("-2,null,-3"), -5, true
  T(_t, S(i1, i2), S(hasPathSum(i1, i2)), S(o))

  i1, i2, o = NewTreeNode(""), 0, false
  T(_t, S(i1, i2), S(hasPathSum(i1, i2)), S(o))

  i1, i2, o = NewTreeNode(""), 1, false
  T(_t, S(i1, i2), S(hasPathSum(i1, i2)), S(o))
}
