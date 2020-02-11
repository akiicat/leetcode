package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestSearchBST(_t *testing.T) {
  i, n, o := NewTreeNode("4,2,7,1,3"), 2, NewTreeNode("2,1,3")
  T(_t, S(i, n), S(searchBST(i, n)), S(o))
}

