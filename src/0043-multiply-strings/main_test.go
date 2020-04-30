package main
import "testing"
import . "main/pkg/testing_helper"

func TestMultiply(_t *testing.T) {
  i1, i2, o := "2", "3", "6"
  T(_t, S(i1, i2), S(multiply(i1, i2)), S(o))

  i1, i2, o = "123", "456", "56088"
  T(_t, S(i1, i2), S(multiply(i1, i2)), S(o))

  i1, i2, o = "0", "0", "0"
  T(_t, S(i1, i2), S(multiply(i1, i2)), S(o))

  i1, i2, o = "9913", "0", "0"
  T(_t, S(i1, i2), S(multiply(i1, i2)), S(o))

  i1, i2, o = "999", "999", "998001"
  T(_t, S(i1, i2), S(multiply(i1, i2)), S(o))
}

func TestMultiplyReverse(_t *testing.T) {
  i1, i2, o := "2", "3", "6"
  T(_t, S(i1, i2), S(multiplyReverse(i1, i2)), S(o))

  i1, i2, o = "123", "456", "56088"
  T(_t, S(i1, i2), S(multiplyReverse(i1, i2)), S(o))

  i1, i2, o = "0", "0", "0"
  T(_t, S(i1, i2), S(multiplyReverse(i1, i2)), S(o))

  i1, i2, o = "9913", "0", "0"
  T(_t, S(i1, i2), S(multiplyReverse(i1, i2)), S(o))

  i1, i2, o = "999", "999", "998001"
  T(_t, S(i1, i2), S(multiplyReverse(i1, i2)), S(o))
}

