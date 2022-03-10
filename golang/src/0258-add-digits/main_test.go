package main
import "testing"
import . "main/pkg/testing_helper"

func TestAddDigits(_t *testing.T) {
  i, o := 38, 2
  T(_t, S(i), S(addDigits(i)), S(o))

  i, o = 0, 0
  T(_t, S(i), S(addDigits(i)), S(o))

  i, o = 18, 9
  T(_t, S(i), S(addDigits(i)), S(o))

  i, o = 9, 9
  T(_t, S(i), S(addDigits(i)), S(o))
}

func TestAddDigitsRecursive(_t *testing.T) {
  i, o := 38, 2
  T(_t, S(i), S(addDigitsRecursive(i)), S(o))

  i, o = 0, 0
  T(_t, S(i), S(addDigits(i)), S(o))

  i, o = 18, 9
  T(_t, S(i), S(addDigits(i)), S(o))

  i, o = 9, 9
  T(_t, S(i), S(addDigits(i)), S(o))
}
