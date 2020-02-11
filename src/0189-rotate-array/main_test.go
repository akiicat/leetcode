package main
import "testing"
import . "main/pkg/testing_helper"

func TestMajorityElement(_t *testing.T) {
  i, k, o := []int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4}
  _i := S(i, k)
  rotate(i, k)
  T(_t, _i, S(i), S(o))

  i, k, o = []int{-1,-100,3,99}, 2, []int{3,99,-1,-100}
  _i = S(i, k)
  rotate(i, k)
  T(_t, _i, S(i), S(o))

  i, k, o = []int{-1}, 2, []int{-1}
  _i = S(i, k)
  rotate(i, k)
  T(_t, _i, S(i), S(o))
}
