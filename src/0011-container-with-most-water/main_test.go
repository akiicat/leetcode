package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxArea(_t *testing.T) {
  i, o := []int{1,8,6,2,5,4,8,3,7}, 49
  T(_t, S(i), S(maxArea(i)), S(o))

  i, o = []int{2,3,4,5,18,17,6}, 17
  T(_t, S(i), S(maxArea(i)), S(o))

  i, o = []int{1,2,3,4,5,25,24,3,4}, 24
  T(_t, S(i), S(maxArea(i)), S(o))
}

func TestMaxAreaDynamicProgramming(_t *testing.T) {
  i, o := []int{1,8,6,2,5,4,8,3,7}, 49
  T(_t, S(i), S(maxAreaDynamicProgramming(i)), S(o))

  i, o = []int{2,3,4,5,18,17,6}, 17
  T(_t, S(i), S(maxAreaDynamicProgramming(i)), S(o))

  i, o = []int{1,2,3,4,5,25,24,3,4}, 24
  T(_t, S(i), S(maxAreaDynamicProgramming(i)), S(o))
}

