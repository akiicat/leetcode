package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsMonotonic(_t *testing.T) {
  i, o := []int{1,2,2,3}, true
  T(_t, S(i), S(isMonotonic(i)), S(o))

  i, o = []int{6,5,4,4}, true
  T(_t, S(i), S(isMonotonic(i)), S(o))

  i, o = []int{1,3,2}, false
  T(_t, S(i), S(isMonotonic(i)), S(o))

  i, o = []int{1,2,4,5}, true
  T(_t, S(i), S(isMonotonic(i)), S(o))

  i, o = []int{1,1,1}, true
  T(_t, S(i), S(isMonotonic(i)), S(o))
}
