package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsPowerOfFour(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfFour(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfFour(i)), S(o))

  i, o = 2, false
  T(_t, S(i), S(isPowerOfFour(i)), S(o))

  i, o = 5, false
  T(_t, S(i), S(isPowerOfFour(i)), S(o))

  i, o = 16, true
  T(_t, S(i), S(isPowerOfFour(i)), S(o))
}

func TestIsPowerOfFourLog(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfFourLog(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfFourLog(i)), S(o))

  i, o = 2, false
  T(_t, S(i), S(isPowerOfFourLog(i)), S(o))

  i, o = 5, false
  T(_t, S(i), S(isPowerOfFourLog(i)), S(o))

  i, o = 16, true
  T(_t, S(i), S(isPowerOfFourLog(i)), S(o))
}

func TestIsPowerOfFourForLoop(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfFourForLoop(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfFourForLoop(i)), S(o))

  i, o = 2, false
  T(_t, S(i), S(isPowerOfFourForLoop(i)), S(o))

  i, o = 5, false
  T(_t, S(i), S(isPowerOfFourForLoop(i)), S(o))

  i, o = 16, true
  T(_t, S(i), S(isPowerOfFourForLoop(i)), S(o))
}

