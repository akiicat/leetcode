package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindSpecialInteger(t *testing.T) {
  i, o := []int{1,2,2,6,6,6,6,7,10}, 6
  T(t, S(i), S(findSpecialInteger(i)), S(o))
}

func TestFindSpecialIntegerLinear(t *testing.T) {
  i, o := []int{1,2,2,6,6,6,6,7,10}, 6
  T(t, S(i), S(findSpecialIntegerLinear(i)), S(o))
}

