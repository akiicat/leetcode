package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaximumProduct(_t *testing.T) {
  i, o := []int{1,2,3}, 6
  T(_t, S(i), S(maximumProduct(i)), S(o))

  i, o = []int{1,2,3,4}, 24
  T(_t, S(i), S(maximumProduct(i)), S(o))

  i, o = []int{-1,2,-3,4}, 12
  T(_t, S(i), S(maximumProduct(i)), S(o))

  i, o = []int{-1,-2,-3}, -6
  T(_t, S(i), S(maximumProduct(i)), S(o))

  i, o = []int{-4,-3,-2,-1,60}, 720
  T(_t, S(i), S(maximumProduct(i)), S(o))
}
