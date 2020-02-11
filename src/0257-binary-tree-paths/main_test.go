package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestIsAnagram(_t *testing.T) {
  i, o := NewTreeNode("1,2,3,null,5"), []string{"1->2->5", "1->3"}
  T(_t, S(i), S(binaryTreePaths(i)), S(o))

  i, o = NewTreeNode("3,9,20,null,null,15,7"), []string{"3->9","3->20->15","3->20->7"}
  T(_t, S(i), S(binaryTreePaths(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), []string{"1->2->3->4","1->2->3->4","1->2->3","1->2"}
  T(_t, S(i), S(binaryTreePaths(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), []string{"1->2->3->4","1->2->3->4"}
  T(_t, S(i), S(binaryTreePaths(i)), S(o))

  i, o = NewTreeNode(""), []string{}
  T(_t, S(i), S(binaryTreePaths(i)), S(o))
}
