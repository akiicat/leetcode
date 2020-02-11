package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestMergeTrees(_t *testing.T) {
  t1, t2, o := NewTreeNode("1,3,2,5"), NewTreeNode("2,1,3,null,4,null,7"), NewTreeNode("3,4,5,5,4,null,7")
  T(_t, S(t1, t2), S(mergeTrees(t1, t2)), S(o))
}

