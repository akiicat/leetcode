package main
import "testing"
import . "main/pkg/testing_helper"

func TestRotatedDigits(_t *testing.T) {
  i, o := 6, 3
  T(_t, S(i), S(rotatedDigits(i)), S(o))

  i, o = 10, 4
  T(_t, S(i), S(rotatedDigits(i)), S(o))

  i, o = 12, 5
  T(_t, S(i), S(rotatedDigits(i)), S(o))

  i, o = 156, 60
  T(_t, S(i), S(rotatedDigits(i)), S(o))

  i, o = 857, 247
  T(_t, S(i), S(rotatedDigits(i)), S(o))
}

func TestRotatedDigitsLinear(_t *testing.T) {
  i, o := 6, 3
  T(_t, S(i), S(rotatedDigitsLinear(i)), S(o))

  i, o = 10, 4
  T(_t, S(i), S(rotatedDigitsLinear(i)), S(o))

  i, o = 12, 5
  T(_t, S(i), S(rotatedDigitsLinear(i)), S(o))

  i, o = 156, 60
  T(_t, S(i), S(rotatedDigitsLinear(i)), S(o))

  i, o = 857, 247
  T(_t, S(i), S(rotatedDigitsLinear(i)), S(o))
}

