package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsBalanced(_t *testing.T) {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), true
  T(_t, S(i), S(isBalanced(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), false
  T(_t, S(i), S(isBalanced(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), false
  T(_t, S(i), S(isBalanced(i)), S(o))

  i, o = NewTreeNode(""), true
  T(_t, S(i), S(isBalanced(i)), S(o))
}


func TestIsBalancedRecursive(_t *testing.T) {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), true
  T(_t, S(i), S(isBalancedRecursive(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), false
  T(_t, S(i), S(isBalancedRecursive(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), false
  T(_t, S(i), S(isBalancedRecursive(i)), S(o))

  i, o = NewTreeNode(""), true
  T(_t, S(i), S(isBalancedRecursive(i)), S(o))
}
