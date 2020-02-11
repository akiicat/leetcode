package main
import "testing"
import . "main/pkg/testing_helper"

func TestLargestPerimeter(t *testing.T) {
  i, o := []int{2,1,2}, 5
  T(t, S(i), S(largestPerimeter(i)), S(o))

  i, o = []int{1,2,1}, 0
  T(t, S(i), S(largestPerimeter(i)), S(o))

  i, o = []int{3,2,3,4}, 10
  T(t, S(i), S(largestPerimeter(i)), S(o))

  i, o = []int{3,6,2,3}, 8
  T(t, S(i), S(largestPerimeter(i)), S(o))
}

