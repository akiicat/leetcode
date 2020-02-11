package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsRectangleOverlap(_t *testing.T) {
  r1, r2, o := []int{0,0,2,2}, []int{1,1,3,3}, true
  T(_t, S(r1, r2), S(isRectangleOverlap(r1, r2)), S(o))

  r1, r2, o = []int{0,0,1,1}, []int{1,0,2,1}, false
  T(_t, S(r1, r2), S(isRectangleOverlap(r1, r2)), S(o))
}

