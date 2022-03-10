package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestTree2str(_t *testing.T) {
  i, o := NewTreeNode("1,2,3,4"), "1(2(4))(3)"
  T(_t, S(i), S(tree2str(i)), S(o))

  i, o = NewTreeNode("1,2,3,null,4"), "1(2()(4))(3)"
  T(_t, S(i), S(tree2str(i)), S(o))
}

