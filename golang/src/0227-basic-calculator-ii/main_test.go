package main
import "testing"
import . "main/pkg/testing_helper"

func TestCalculate(_t *testing.T) {
  i, o := "3+2*2", 7
  T(_t, S(i), S(calculate(i)), S(o))

  i, o = " 3/2 ", 1
  T(_t, S(i), S(calculate(i)), S(o))

  i, o = " 3+5 / 2 ", 5
  T(_t, S(i), S(calculate(i)), S(o))

  i, o = "2*2/4", 1
  T(_t, S(i), S(calculate(i)), S(o))

  i, o = "2/4*2", 0
  T(_t, S(i), S(calculate(i)), S(o))

  i, o = "1-1", 0
  T(_t, S(i), S(calculate(i)), S(o))
}

