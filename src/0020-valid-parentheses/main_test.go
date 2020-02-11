package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsValid(_t *testing.T) {
  i, o := "()", true
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "()", true
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "()[]{}", true
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "{[]}", true
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "[", false
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "]", false
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "(]", false
  T(_t, S(i), S(isValid(i)), S(o))

  i, o = "([)]", false
  T(_t, S(i), S(isValid(i)), S(o))
}

