package main
import "testing"
import . "main/pkg/testing_helper"

func TestLargeGroupPositions(_t *testing.T) {
  i, o := "abbxxxxzyy", [][]int{[]int{3,6}}
  T(_t, S(i), S(largeGroupPositions(i)), S(o))

  i, o = "abc", [][]int{}
  T(_t, S(i), S(largeGroupPositions(i)), S(o))

  i, o = "abcdddeeeeaabbbcd", [][]int{[]int{3,5},[]int{6,9},[]int{12,14}}
  T(_t, S(i), S(largeGroupPositions(i)), S(o))

  i, o = "aaa", [][]int{[]int{0,2}}
  T(_t, S(i), S(largeGroupPositions(i)), S(o))
}

