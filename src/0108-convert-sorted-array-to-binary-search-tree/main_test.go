package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestSortedArrayToBST(_t *testing.T) {
  i, o := []int{-10,-3,0,5,9}, NewTreeNode("0,-3,9,-10,null,5")
  T(_t, S(i), S(sortedArrayToBST(i)), S(o))
}

