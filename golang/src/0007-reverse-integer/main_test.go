package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverse(_t *testing.T) {
  i, o := 123, 321
  T(_t, S(i), S(reverse(i)), S(o))

  i, o = -123, -321
  T(_t, S(i), S(reverse(i)), S(o))

  i, o = 120, 21
  T(_t, S(i), S(reverse(i)), S(o))

  i, o = 1534236469, 0
  T(_t, S(i), S(reverse(i)), S(o))
}

