package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestSumOfLeftLeaves(_t *testing.T) {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), 24
  T(_t, S(i), S(sumOfLeftLeaves(i)), S(o))

  i, o = NewTreeNode("1"), 0
  T(_t, S(i), S(sumOfLeftLeaves(i)), S(o))
}

func TestSumOfLeftLeavesStack(_t *testing.T) {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), 24
  T(_t, S(i), S(sumOfLeftLeavesStack(i)), S(o))

  i, o = NewTreeNode("1"), 0
  T(_t, S(i), S(sumOfLeftLeavesStack(i)), S(o))
}
