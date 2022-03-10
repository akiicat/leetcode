package main
import "testing"
import . "main/pkg/testing_helper"

func TestSmallestRangeI(t *testing.T) {
  i, k, o := []int{1}, 0, 0
  T(t, S(i, k), S(smallestRangeI(i, k)), S(o))

  i, k, o = []int{0,10}, 0, 10
  T(t, S(i, k), S(smallestRangeI(i, k)), S(o))

  i, k, o = []int{0,10}, 2, 6
  T(t, S(i, k), S(smallestRangeI(i, k)), S(o))

  i, k, o = []int{0,10}, 10, 0
  T(t, S(i, k), S(smallestRangeI(i, k)), S(o))

  i, k, o = []int{1,3,6}, 3, 0
  T(t, S(i, k), S(smallestRangeI(i, k)), S(o))
}

