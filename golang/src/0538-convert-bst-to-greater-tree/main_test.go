package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestConvertBST(_t *testing.T) {
  i, o := NewTreeNode("5,2,13"), NewTreeNode("18,20,13")
  T(_t, S(i), S(convertBST(i)), S(o))
}

