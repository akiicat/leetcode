package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindPairs(_t *testing.T) {
  i, k, o := []int{3,1,4,1,5}, 2, 2
  T(_t, S(i, k), S(findPairs(i, k)), S(o))

  i, k, o = []int{1,2,3,4,5}, 1, 4
  T(_t, S(i, k), S(findPairs(i, k)), S(o))

  i, k, o = []int{1,3,1,5,4}, 0, 1
  T(_t, S(i, k), S(findPairs(i, k)), S(o))

  i, k, o = []int{1,1,1,1,1}, 0, 1
  T(_t, S(i, k), S(findPairs(i, k)), S(o))

  i, k, o = []int{-1,0,1,-2,0}, 2, 2
  T(_t, S(i, k), S(findPairs(i, k)), S(o))

  i, k, o = []int{-1,0,1,-2,0}, -1, 0
  T(_t, S(i, k), S(findPairs(i, k)), S(o))
}
