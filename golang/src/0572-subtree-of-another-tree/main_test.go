package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsSubtree(_t *testing.T) {
  s, t, o := NewTreeNode("3,4,5,1,2"), NewTreeNode(""), false
  T(_t, S(s, t), S(isSubtree(s, t)), S(o))

  s, t, o = NewTreeNode("3,4,5,1,2"), NewTreeNode("4,1,2"), true
  T(_t, S(s, t), S(isSubtree(s, t)), S(o))

  s, t, o = NewTreeNode("3,4,5,1,2,null,null,0"), NewTreeNode("4,1,2"), false
  T(_t, S(s, t), S(isSubtree(s, t)), S(o))

  s, t, o = NewTreeNode("1,1"), NewTreeNode("1"), true
  T(_t, S(s, t), S(isSubtree(s, t)), S(o))
}

