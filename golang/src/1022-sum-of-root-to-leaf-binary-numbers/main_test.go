package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestSumRootToLeaf(_t *testing.T) {
  i, o := NewTreeNode("1,0,1,0,1,0,1"), 22
  T(_t, S(i), S(sumRootToLeaf(i)), S(o))

  i, o = NewTreeNode("0,1,0,0,null,0,0,null,null,null,1,null,null,null,1"), 5
  T(_t, S(i), S(sumRootToLeaf(i)), S(o))
}

