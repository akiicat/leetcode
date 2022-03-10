package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestMinDiffInBST(_t *testing.T) {
  i, o := NewTreeNode("4,2,6,1,3,null,null"), 1
  T(_t, S(i), S(minDiffInBST(i)), S(o))
}

func TestMinDiffInBSTDFS(_t *testing.T) {
  i, o := NewTreeNode("4,2,6,1,3,null,null"), 1
  T(_t, S(i), S(minDiffInBSTDFS(i)), S(o))
}

