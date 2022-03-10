package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestPathSum(_t *testing.T) {
  i, s, o := NewTreeNode("10,5,-3,3,2,null,11,3,-2,null,1"), 8, 3
  T(_t, S(i, s), S(pathSum(i, s)), S(o))

  i, s, o = NewTreeNode("1,-2,-3"), -1, 1
  T(_t, S(i, s), S(pathSum(i, s)), S(o))
}

func TestPathSumRecursive(_t *testing.T) {
  i, s, o := NewTreeNode("10,5,-3,3,2,null,11,3,-2,null,1"), 8, 3
  T(_t, S(i, s), S(pathSumRecursive(i, s)), S(o))

  i, s, o = NewTreeNode("1,-2,-3"), -1, 1
  T(_t, S(i, s), S(pathSumRecursive(i, s)), S(o))
}

func TestPathSumArray(_t *testing.T) {
  i, s, o := NewTreeNode("10,5,-3,3,2,null,11,3,-2,null,1"), 8, 3
  T(_t, S(i, s), S(pathSumArray(i, s)), S(o))

  i, s, o = NewTreeNode("1,-2,-3"), -1, 1
  T(_t, S(i, s), S(pathSumArray(i, s)), S(o))
}
