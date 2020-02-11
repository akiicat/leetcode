package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindLengthOfLCIS(_t *testing.T) {
  i, o := []int{1,3,5,4,7}, 3
  T(_t, S(i), S(findLengthOfLCIS(i)), S(o))

  i, o = []int{2,2,2,2,2}, 1
  T(_t, S(i), S(findLengthOfLCIS(i)), S(o))

  i, o = []int{1,3,5,7}, 4
  T(_t, S(i), S(findLengthOfLCIS(i)), S(o))

  i, o = []int{}, 0
  T(_t, S(i), S(findLengthOfLCIS(i)), S(o))
}

