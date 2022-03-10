package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestMinDepth(_t *testing.T) {
  i, o := NewTreeNode("1,2,3,null,5"), 2
  T(_t, S(i), S(minDepth(i)), S(o))

  i, o = NewTreeNode("3,9,20,null,null,15,7"), 2
  T(_t, S(i), S(minDepth(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), 2
  T(_t, S(i), S(minDepth(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), 4
  T(_t, S(i), S(minDepth(i)), S(o))

  i, o = NewTreeNode("1,2"), 2
  T(_t, S(i), S(minDepth(i)), S(o))

  i, o = NewTreeNode("1"), 1
  T(_t, S(i), S(minDepth(i)), S(o))

  i, o = NewTreeNode(""), 0
  T(_t, S(i), S(minDepth(i)), S(o))
}
