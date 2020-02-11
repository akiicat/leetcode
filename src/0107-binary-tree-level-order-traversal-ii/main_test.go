package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestLevelOrderBottom(_t *testing.T) {
  i, o := NewTreeNode("1,2,3,null,5"), [][]int{[]int{5}, []int{2,3}, []int{1}}
  T(_t, S(i), S(levelOrderBottom(i)), S(o))

  i, o = NewTreeNode("1,2,3,4,5"), [][]int{[]int{4,5}, []int{2,3}, []int{1}}
  T(_t, S(i), S(levelOrderBottom(i)), S(o))

  i, o = NewTreeNode("3,9,20,null,null,15,7"), [][]int{[]int{15,7}, []int{9,20}, []int{3}}
  T(_t, S(i), S(levelOrderBottom(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), [][]int{[]int{4,4},[]int{3,3},[]int{2,2},[]int{1}}
  T(_t, S(i), S(levelOrderBottom(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), [][]int{[]int{4,4},[]int{3,3},[]int{2,2},[]int{1}}
  T(_t, S(i), S(levelOrderBottom(i)), S(o))

  i, o = NewTreeNode(""), [][]int{}
  T(_t, S(i), S(levelOrderBottom(i)), S(o))
}

func TestLevelOrderBottomRecursive(_t *testing.T) {
  i, o := NewTreeNode("1,2,3,null,5"), [][]int{[]int{5}, []int{2,3}, []int{1}}
  T(_t, S(i), S(levelOrderBottomRecursive(i)), S(o))

  i, o = NewTreeNode("1,2,3,4,5"), [][]int{[]int{4,5}, []int{2,3}, []int{1}}
  T(_t, S(i), S(levelOrderBottomRecursive(i)), S(o))

  i, o = NewTreeNode("3,9,20,null,null,15,7"), [][]int{[]int{15,7}, []int{9,20}, []int{3}}
  T(_t, S(i), S(levelOrderBottomRecursive(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), [][]int{[]int{4,4},[]int{3,3},[]int{2,2},[]int{1}}
  T(_t, S(i), S(levelOrderBottomRecursive(i)), S(o))

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), [][]int{[]int{4,4},[]int{3,3},[]int{2,2},[]int{1}}
  T(_t, S(i), S(levelOrderBottomRecursive(i)), S(o))

  i, o = NewTreeNode(""), [][]int{}
  T(_t, S(i), S(levelOrderBottomRecursive(i)), S(o))
}

