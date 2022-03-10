package main
import "testing"
import . "main/pkg/testing_helper"

func TestTwoSum(_t *testing.T) {
  i, t, o := []int{2,7,11,19}, 9, []int{1,2}
  T(_t, S(i, t), S(twoSum(i, t)), S(o))

  i, t, o = []int{2,3,4}, 6, []int{1,3}
  T(_t, S(i, t), S(twoSum(i, t)), S(o))
}
