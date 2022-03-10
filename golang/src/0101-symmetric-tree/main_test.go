package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsSymmetric(_t *testing.T) {
  //     1
  //    / \
  //   2   2
  //  / \ / \
  // 3  4 4  3
  i, o := NewTreeNode("1,2,2,3,4,4,3"), true
  T(_t, S(i), S(isSymmetric(i)), S(o))

  //   1
  //  / \
  // 2   2
  //  \   \
  //  3    3
  i, o = NewTreeNode("1,2,2,null,3,null,3"), false
  T(_t, S(i), S(isSymmetric(i)), S(o))

  i, o = NewTreeNode(""), true
  T(_t, S(i), S(isSymmetric(i)), S(o))
}

func TestIsSymmetricIterative(_t *testing.T) {
  i, o := NewTreeNode("1,2,2,3,4,4,3"), true
  T(_t, S(i), S(isSymmetricIterative(i)), S(o))

  i, o = NewTreeNode("1,2,2,null,3,null,3"), false
  T(_t, S(i), S(isSymmetricIterative(i)), S(o))

  i, o = NewTreeNode(""), true
  T(_t, S(i), S(isSymmetricIterative(i)), S(o))
}

