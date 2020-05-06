package main
import "testing"
import . "main/pkg/testing_helper"

func TestCompress(_t *testing.T) {
  i, o := "42", 42
  T(_t, S(i), S(myAtoi(i)), S(o))

  i, o = "  -42", -42
  T(_t, S(i), S(myAtoi(i)), S(o))

  i, o = "4193 with words", 4193
  T(_t, S(i), S(myAtoi(i)), S(o))

  i, o = "words and 987", 0
  T(_t, S(i), S(myAtoi(i)), S(o))

  i, o = "-91283472332", -2147483648
  T(_t, S(i), S(myAtoi(i)), S(o))

  i, o = "9223372036854775808", 2147483647
  T(_t, S(i), S(myAtoi(i)), S(o))
}

