package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindContentChildren(_t *testing.T) {
  i1, i2, o := []int{1,2,3}, []int{1,1}, 1
  T(_t, S(i1, i2), S(findContentChildren(i1, i2)), S(o))

  i1, i2, o = []int{1,2}, []int{1,2,3}, 2
  T(_t, S(i1, i2), S(findContentChildren(i1, i2)), S(o))

  i1, i2, o = []int{1,2}, []int{1,3}, 2
  T(_t, S(i1, i2), S(findContentChildren(i1, i2)), S(o))

  i1, i2, o = []int{1,2,3}, []int{}, 0
  T(_t, S(i1, i2), S(findContentChildren(i1, i2)), S(o))

  i1, i2, o = []int{1,2,3}, []int{3}, 1
  T(_t, S(i1, i2), S(findContentChildren(i1, i2)), S(o))
}


