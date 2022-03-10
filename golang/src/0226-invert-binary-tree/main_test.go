package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestInvertTree(_t *testing.T) {
  i, o := NewTreeNode("4,2,7,1,3,6,9"), NewTreeNode("4,7,2,9,6,3,1")
  T(_t, S(i), S(invertTree(i)), S(o))
}

