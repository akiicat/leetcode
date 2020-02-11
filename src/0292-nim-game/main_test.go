package main
import "testing"
import . "main/pkg/testing_helper"

func TestCanWinNim(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 2, true
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 3, true
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 4, false
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 6, true
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 7, true
  T(_t, S(i), S(canWinNim(i)), S(o))

  i, o = 8, false
  T(_t, S(i), S(canWinNim(i)), S(o))
}
