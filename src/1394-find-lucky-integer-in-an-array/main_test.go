package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindLucky(_t *testing.T) {
  i, o := []int{2,2,3,4}, 2
  T(_t, S(i), S(findLucky(i)), S(o))

  i, o = []int{1,2,2,3,3,3}, 3
  T(_t, S(i), S(findLucky(i)), S(o))

  i, o = []int{2,2,2,3,3}, -1
  T(_t, S(i), S(findLucky(i)), S(o))

  i, o = []int{5}, -1
  T(_t, S(i), S(findLucky(i)), S(o))

  i, o = []int{7,7,7,7,7,7,7}, 7
  T(_t, S(i), S(findLucky(i)), S(o))
}

