package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestLeafSimilar(_t *testing.T) {
  i1, i2, o := NewTreeNode("3,5,1,6,2,9,8,null,null,7,4"), NewTreeNode("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"), true
  T(_t, S(i1, i2), S(leafSimilar(i1, i2)), S(o))
}

