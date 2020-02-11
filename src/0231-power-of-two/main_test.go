package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsPowerOfTwo(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))

  i, o = 2, true
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))

  i, o = 4, true
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))

  i, o = 16, true
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))

  i, o = 218, false
  T(_t, S(i), S(isPowerOfTwo(i)), S(o))
}
