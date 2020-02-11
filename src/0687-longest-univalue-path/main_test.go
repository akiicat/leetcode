package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestLongestUnivaluePath(_t *testing.T) {
  i, o := NewTreeNode("5,4,5,1,1,5"), 2
  T(_t, S(i), S(longestUnivaluePath(i)), S(o))

  i, o = NewTreeNode("1,4,5,4,4,5"), 2
  T(_t, S(i), S(longestUnivaluePath(i)), S(o))
}

