package main
import "testing"
import . "main/pkg/testing_helper"

func TestTrailingZeroes(_t *testing.T) {
  i, o := 3, 0
  T(_t, S(i), S(trailingZeroes(i)), S(o))

  i, o = 5, 1
  T(_t, S(i), S(trailingZeroes(i)), S(o))

  i, o = 25, 6
  T(_t, S(i), S(trailingZeroes(i)), S(o))
}
