package main
import "testing"
import . "main/pkg/testing_helper"

func TestToHex(_t *testing.T) {
  i, o := 26, "1a"
  T(_t, S(i), S(toHex(i)), S(o))

  i, o = 25, "19"
  T(_t, S(i), S(toHex(i)), S(o))

  i, o = -1, "ffffffff"
  T(_t, S(i), S(toHex(i)), S(o))
}
