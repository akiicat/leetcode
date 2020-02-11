package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestFindMode(_t *testing.T) {
  i, o := NewTreeNode("1,null,2,2"), []int{2}
  T(_t, S(i), S(findMode(i)), S(o))

  i, o = NewTreeNode("2147483647"), []int{2147483647}
  T(_t, S(i), S(findMode(i)), S(o))
}

func TestFindModeMap(_t *testing.T) {
  i, o := NewTreeNode("1,null,2,2"), []int{2}
  T(_t, S(i), S(findModeMap(i)), S(o))

  i, o = NewTreeNode("2147483647"), []int{2147483647}
  T(_t, S(i), S(findModeMap(i)), S(o))
}

func TestFindModeArray(_t *testing.T) {
  i, o := NewTreeNode("1,null,2,2"), []int{2}
  T(_t, S(i), S(findModeArray(i)), S(o))

  i, o = NewTreeNode("2147483647"), []int{2147483647}
  T(_t, S(i), S(findModeArray(i)), S(o))
}
