package main
import "testing"
import . "main/pkg/testing_helper"

func TestNextGreaterElement(_t *testing.T) {
  i1, i2, o := []int{4,1,2}, []int{1,3,4,2}, []int{-1,3,-1}
  T(_t, S(i1, i2), S(nextGreaterElement(i1, i2)), S(o))

  i1, i2, o = []int{2,4}, []int{1,2,3,4}, []int{3,-1}
  T(_t, S(i1, i2), S(nextGreaterElement(i1, i2)), S(o))

  i1, i2, o = []int{3,1,5,7,9,2,6}, []int{1,2,3,5,6,7,9,11}, []int{5,2,6,9,11,3,7}
  T(_t, S(i1, i2), S(nextGreaterElement(i1, i2)), S(o))
}

