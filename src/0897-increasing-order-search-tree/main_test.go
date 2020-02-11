package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIncreasingBST(t *testing.T) {
  i, o := NewTreeNode("5,3,6,2,4,null,8,1,null,null,null,7,9"), NewTreeNode("1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9")
  T(t, S(i), S(increasingBST(i)), S(o))
}

func TestIncreasingBSTReverseRecursive(t *testing.T) {
  i, o := NewTreeNode("5,3,6,2,4,null,8,1,null,null,null,7,9"), NewTreeNode("1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9")
  T(t, S(i), S(increasingBSTReverseRecursive(i)), S(o))
}

