package main
import "testing"
import . "main/pkg/testing_helper"

func TestLengthOfLastWord(_t *testing.T) {
  i, o := "Hello World", 5
  T(_t, S(i), S(lengthOfLastWord(i)), S(o))

  i, o = "Hello ", 5
  T(_t, S(i), S(lengthOfLastWord(i)), S(o))

  i, o = "  ", 0
  T(_t, S(i), S(lengthOfLastWord(i)), S(o))

  i, o = " ", 0
  T(_t, S(i), S(lengthOfLastWord(i)), S(o))

  i, o = "", 0
  T(_t, S(i), S(lengthOfLastWord(i)), S(o))
}

func TestLengthOfLastWordBoolean(_t *testing.T) {
  i, o := "Hello World", 5
  T(_t, S(i), S(lengthOfLastWordBoolean(i)), S(o))

  i, o = "Hello ", 5
  T(_t, S(i), S(lengthOfLastWordBoolean(i)), S(o))

  i, o = "  ", 0
  T(_t, S(i), S(lengthOfLastWordBoolean(i)), S(o))

  i, o = " ", 0
  T(_t, S(i), S(lengthOfLastWordBoolean(i)), S(o))

  i, o = "", 0
  T(_t, S(i), S(lengthOfLastWordBoolean(i)), S(o))
}
