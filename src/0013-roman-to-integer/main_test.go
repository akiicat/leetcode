package main
import "testing"
import . "main/pkg/testing_helper"

func TestRomanToInt(_t *testing.T) {
  i, o := "III", 3
  T(_t, S(i), S(romanToInt(i)), S(o))

  i, o = "IV", 4
  T(_t, S(i), S(romanToInt(i)), S(o))

  i, o = "IX", 9
  T(_t, S(i), S(romanToInt(i)), S(o))

  i, o = "LVIII", 58
  T(_t, S(i), S(romanToInt(i)), S(o))

  i, o = "MCMXCIV", 1994
  T(_t, S(i), S(romanToInt(i)), S(o))
}
