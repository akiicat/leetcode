package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestMaxDepth(_t *testing.T) {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), 3
  T(_t, S(i), S(maxDepth(i)), S(o))
}

