package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsUgly(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isUgly(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isUgly(i)), S(o))

  i, o = 6, true
  T(_t, S(i), S(isUgly(i)), S(o))

  i, o = 8, true
  T(_t, S(i), S(isUgly(i)), S(o))

  i, o = 14, false
  T(_t, S(i), S(isUgly(i)), S(o))
}
