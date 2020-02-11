package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestRangeSumBST(t *testing.T) {
  // 10,5,15,3,7,null,18
  //                10
  //              /    \
  //           5         15
  //          / \      /    \
  //        3    7    null  18
  i, L, R, o := NewTreeNode("10,5,15,3,7,null,18"), 7, 15, 32
  T(t, S(i, L, R), S(rangeSumBST(i, L, R)), S(o))

  // 10,5,15,3,7,13,18,1,null,6
  //                  10
  //              /       \
  //             5        15
  //         /      \    /   \
  //        3        7 13     18
  //       / \      /
  //      1  null  6
  i, L, R, o = NewTreeNode("10,5,15,3,7,13,18,1,null,6"), 6, 10, 23
  T(t, S(i, L, R), S(rangeSumBST(i, L, R)), S(o))
}

func TestRangeSumBSTMergeAlgorithm(t *testing.T) {
  i, L, R, o := NewTreeNode("10,5,15,3,7,null,18"), 7, 15, 32
  T(t, S(i, L, R), S(rangeSumBSTMergeAlgorithm(i, L, R)), S(o))

  i, L, R, o = NewTreeNode("10,5,15,3,7,13,18,1,null,6"), 6, 10, 23
  T(t, S(i, L, R), S(rangeSumBSTMergeAlgorithm(i, L, R)), S(o))
}

func TestRangeSumBSTDFSButtomUpRecursive(t *testing.T) {
  i, L, R, o := NewTreeNode("10,5,15,3,7,null,18"), 7, 15, 32
  T(t, S(i, L, R), S(rangeSumBSTDFSButtomUpRecursive(i, L, R)), S(o))

  i, L, R, o = NewTreeNode("10,5,15,3,7,13,18,1,null,6"), 6, 10, 23
  T(t, S(i, L, R), S(rangeSumBSTDFSButtomUpRecursive(i, L, R)), S(o))
}

