package main
import "testing"
import . "main/pkg/testing_helper"

func TestCreateTargetArray(_t *testing.T) {
  nums, index, o := []int{0,1,2,3,4}, []int{0,1,2,2,1}, []int{0,4,1,3,2}
  T(_t, S(nums, index), S(createTargetArray(nums, index)), S(o))

  nums, index, o = []int{1,2,3,4,0}, []int{0,1,2,3,0}, []int{0,1,2,3,4}
  T(_t, S(nums, index), S(createTargetArray(nums, index)), S(o))

  nums, index, o = []int{1}, []int{0}, []int{1}
  T(_t, S(nums, index), S(createTargetArray(nums, index)), S(o))
}

