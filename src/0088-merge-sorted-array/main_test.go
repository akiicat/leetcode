package main
import "testing"
import . "main/pkg/testing_helper"

func TestMerge(_t *testing.T) {
  i1, i2, i3, i4 := []int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3
  o := []int{1,2,2,3,5,6}
  i := S(i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  T(_t, i, S(i1), S(o))

  i1, i2, i3, i4 = []int{1,2,4,5,6,0}, 5, []int{3}, 1
  o = []int{1,2,3,4,5,6}
  i = S(i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  T(_t, i, S(i1), S(o))

  i1, i2, i3, i4 = []int{1}, 1, []int{}, 0
  o = []int{1}
  i = S(i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  T(_t, i, S(i1), S(o))

  i1, i2, i3, i4 = []int{0}, 0, []int{1}, 1
  o = []int{1}
  i = S(i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  T(_t, i, S(i1), S(o))
}
