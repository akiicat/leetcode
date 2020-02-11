package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsPowerOfThree(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfThree(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfThree(i)), S(o))

  i, o = 3, true
  T(_t, S(i), S(isPowerOfThree(i)), S(o))

  i, o = 9, true
  T(_t, S(i), S(isPowerOfThree(i)), S(o))

  i, o = 27, true
  T(_t, S(i), S(isPowerOfThree(i)), S(o))

  i, o = 45, false
  T(_t, S(i), S(isPowerOfThree(i)), S(o))
}

func TestIsPowerOfThreeLog(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfThreeLog(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfThreeLog(i)), S(o))

  i, o = 3, true
  T(_t, S(i), S(isPowerOfThreeLog(i)), S(o))

  i, o = 9, true
  T(_t, S(i), S(isPowerOfThreeLog(i)), S(o))

  i, o = 27, true
  T(_t, S(i), S(isPowerOfThreeLog(i)), S(o))

  i, o = 45, false
  T(_t, S(i), S(isPowerOfThreeLog(i)), S(o))
}

func TestIsPowerOfThreeForLoop(_t *testing.T) {
  i, o := 0, false
  T(_t, S(i), S(isPowerOfThreeForLoop(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPowerOfThreeForLoop(i)), S(o))

  i, o = 3, true
  T(_t, S(i), S(isPowerOfThreeForLoop(i)), S(o))

  i, o = 9, true
  T(_t, S(i), S(isPowerOfThreeForLoop(i)), S(o))

  i, o = 27, true
  T(_t, S(i), S(isPowerOfThreeForLoop(i)), S(o))

  i, o = 45, false
  T(_t, S(i), S(isPowerOfThreeForLoop(i)), S(o))
}
