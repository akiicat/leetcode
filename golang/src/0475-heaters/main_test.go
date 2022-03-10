package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindRadius(_t *testing.T) {
  i1, i2, o := []int{1,2,3}, []int{2}, 1
  T(_t, S(i1, i2), S(findRadius(i1, i2)), S(o))

  i1, i2, o = []int{1,2,3,4}, []int{1,4}, 1
  T(_t, S(i1, i2), S(findRadius(i1, i2)), S(o))
}

func TestFindRadiusBruteFroce(_t *testing.T) {
  i1, i2, o := []int{1,2,3}, []int{2}, 1
  T(_t, S(i1, i2), S(findRadiusBruteFroce(i1, i2)), S(o))

  i1, i2, o = []int{1,2,3,4}, []int{1,4}, 1
  T(_t, S(i1, i2), S(findRadiusBruteFroce(i1, i2)), S(o))
}
