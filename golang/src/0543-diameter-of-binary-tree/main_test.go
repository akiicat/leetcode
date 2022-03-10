package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestDiameterOfBinaryTree(_t *testing.T) {
  i, o := NewTreeNode("1,2,3,4,5"), 3
  T(_t, S(i), S(diameterOfBinaryTree(i)), S(o))

  i, o = NewTreeNode(""), 0
  T(_t, S(i), S(diameterOfBinaryTree(i)), S(o))

  i, o = NewTreeNode("4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2"), 8
  T(_t, S(i), S(diameterOfBinaryTree(i)), S(o))
}

