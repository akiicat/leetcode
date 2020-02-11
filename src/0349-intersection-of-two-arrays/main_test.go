package main
import "testing"
import . "main/pkg/testing_helper"

func TestIntersection(_t *testing.T) {
  i1, i2, o := []int{1,2,2,1}, []int{2,2}, []int{2}
  T(_t, S(i1, i2), S(intersection(i1, i2)), S(o))

  i1, i2, o = []int{4,9,5}, []int{9,4,9,8,4}, []int{9,4}
  T(_t, S(i1, i2), S(intersection(i1, i2)), S(o))
}

func TestIntersectionDoubleArray(_t *testing.T) {
  i1, i2, o := []int{1,2,2,1}, []int{2,2}, []int{2}
  T(_t, S(i1, i2), S(intersectionDoubleArray(i1, i2)), S(o))

  i1, i2, o = []int{4,9,5}, []int{9,4,9,8,4}, []int{4,9}
  T(_t, S(i1, i2), S(intersectionDoubleArray(i1, i2)), S(o))
}
