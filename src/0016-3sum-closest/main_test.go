package main
import "testing"
import . "main/pkg/testing_helper"

func TestThreeSumClosest(_t *testing.T) {
  i1, i2, o := []int{-1,2,1,-4}, 1, 2
  T(_t, S(i1, i2), S(threeSumClosest(i1, i2)), S(o))

  i1, i2, o = []int{0,2,1,-3}, 1, 0
  T(_t, S(i1, i2), S(threeSumClosest(i1, i2)), S(o))
}

