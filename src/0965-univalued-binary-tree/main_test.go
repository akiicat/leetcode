package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsUnivalTree(_t *testing.T) {
  i, o := NewTreeNode("1,1,1,1,1,null,1"), true
  T(_t, S(i), S(isUnivalTree(i)), S(o))

  i, o = NewTreeNode("2,2,2,5,2"), false
  T(_t, S(i), S(isUnivalTree(i)), S(o))

  i, o = NewTreeNode("2,2,2,2,5"), false
  T(_t, S(i), S(isUnivalTree(i)), S(o))
}

