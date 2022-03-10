package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestLowestCommonAncestor(_t *testing.T) {
  i, o := NewTreeNode("6,2,8,0,4,7,9,null,null,3,5"), 6
  p, q := i.Left, i.Right
  T(_t, S(i, p.Val, q.Val), S(lowestCommonAncestor(i, p, q).Val), S(o))
}

func TestLowestCommonAncestorArray(_t *testing.T) {
  i, o := NewTreeNode("6,2,8,0,4,7,9,null,null,3,5"), 6
  p, q := i.Left, i.Right
  T(_t, S(i, p.Val, q.Val), S(lowestCommonAncestorArray(i, p, q).Val), S(o))
}

