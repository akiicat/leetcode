package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinStartValue(_t *testing.T) {
  i, o := []int{-3,2,-3,4,2}, 5
  T(_t, S(i), S(minStartValue(i)), S(o))

  i, o = []int{1,2}, 1
  T(_t, S(i), S(minStartValue(i)), S(o))

  i, o = []int{1,-2,-3}, 5
  T(_t, S(i), S(minStartValue(i)), S(o))
}

