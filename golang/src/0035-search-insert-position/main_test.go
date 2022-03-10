package main
import "testing"
import . "main/pkg/testing_helper"

func TestSearchInsert(_t *testing.T) {
  i, n, o := []int{1,3,5,6}, 5, 2
  T(_t, S(i, n), S(searchInsert(i, n)), S(o))

  i, n, o = []int{1,3,5,6}, 2, 1
  T(_t, S(i, n), S(searchInsert(i, n)), S(o))

  i, n, o = []int{1,3,5,6}, 7, 4
  T(_t, S(i, n), S(searchInsert(i, n)), S(o))

  i, n, o = []int{1,3,5,6}, 0, 0
  T(_t, S(i, n), S(searchInsert(i, n)), S(o))
}
