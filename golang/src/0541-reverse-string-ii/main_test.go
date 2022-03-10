package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverseStr(_t *testing.T) {
  i, k, o := "abcdefg", 2, "bacdfeg"
  T(_t, S(i, k), S(reverseStr(i, k)), S(o))

  i, k, o = "a", 2, "a"
  T(_t, S(i, k), S(reverseStr(i, k)), S(o))

  i, k, o = "abcdefg", 3, "cbadefg"
  T(_t, S(i, k), S(reverseStr(i, k)), S(o))
}
