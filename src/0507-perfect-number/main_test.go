package main
import "testing"
import . "main/pkg/testing_helper"

func TestCheckPerfectNumber(_t *testing.T) {
  i, o := 28, true
  T(_t, S(i), S(checkPerfectNumber(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(checkPerfectNumber(i)), S(o))

  i, o = 0, false
  T(_t, S(i), S(checkPerfectNumber(i)), S(o))
}

func TestCheckPerfectNumberSqrt(_t *testing.T) {
  i, o := 28, true
  T(_t, S(i), S(checkPerfectNumberSqrt(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(checkPerfectNumberSqrt(i)), S(o))

  i, o = 0, false
  T(_t, S(i), S(checkPerfectNumberSqrt(i)), S(o))
}

func TestCheckPerfectNumberLinear(_t *testing.T) {
  i, o := 28, true
  T(_t, S(i), S(checkPerfectNumberLinear(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(checkPerfectNumberLinear(i)), S(o))

  i, o = 0, false
  T(_t, S(i), S(checkPerfectNumberLinear(i)), S(o))
}

func TestCheckPerfectNumberOneLine(_t *testing.T) {
  i, o := 28, true
  T(_t, S(i), S(checkPerfectNumberOneLine(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(checkPerfectNumberOneLine(i)), S(o))

  i, o = 0, false
  T(_t, S(i), S(checkPerfectNumberOneLine(i)), S(o))
}

