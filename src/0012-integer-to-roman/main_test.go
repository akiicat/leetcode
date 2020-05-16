package main
import "testing"
import . "main/pkg/testing_helper"

func TestIntToRoman(_t *testing.T) {
  i, o := 3, "III"
  T(_t, S(i), S(intToRoman(i)), S(o))

  i, o = 4, "IV"
  T(_t, S(i), S(intToRoman(i)), S(o))

  i, o = 9, "IX"
  T(_t, S(i), S(intToRoman(i)), S(o))

  i, o = 58, "LVIII"
  T(_t, S(i), S(intToRoman(i)), S(o))

  i, o = 1994, "MCMXCIV"
  T(_t, S(i), S(intToRoman(i)), S(o))
}

func TestIntToRomanMap(_t *testing.T) {
  i, o := 3, "III"
  T(_t, S(i), S(intToRomanMap(i)), S(o))

  i, o = 4, "IV"
  T(_t, S(i), S(intToRomanMap(i)), S(o))

  i, o = 9, "IX"
  T(_t, S(i), S(intToRomanMap(i)), S(o))

  i, o = 58, "LVIII"
  T(_t, S(i), S(intToRomanMap(i)), S(o))

  i, o = 1994, "MCMXCIV"
  T(_t, S(i), S(intToRomanMap(i)), S(o))
}

