package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsSameTree(_t *testing.T) {
  i1, i2, o := NewTreeNode("1,2,3"), NewTreeNode("1,null,3"), false
  T(_t, S(i1, i2), S(isSameTree(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("1,2,3"), NewTreeNode("1,2,3"), true
  T(_t, S(i1, i2), S(isSameTree(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("10,5,15"), NewTreeNode("10,5,null,null,15"), false
  T(_t, S(i1, i2), S(isSameTree(i1, i2)), S(o))
}

func TestIsSameTreeConcise(_t *testing.T) {
  i1, i2, o := NewTreeNode("1,2,3"), NewTreeNode("1,null,3"), false
  T(_t, S(i1, i2), S(isSameTreeConcise(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("1,2,3"), NewTreeNode("1,2,3"), true
  T(_t, S(i1, i2), S(isSameTreeConcise(i1, i2)), S(o))

  i1, i2, o = NewTreeNode("10,5,15"), NewTreeNode("10,5,null,null,15"), false
  T(_t, S(i1, i2), S(isSameTreeConcise(i1, i2)), S(o))
}
