package main
import "testing"
import . "main/pkg/testing_helper"

func TestConvertToBase7(_t *testing.T) {
  i, o := 100, "202"
  T(_t, S(i), S(convertToBase7(i)), S(o))

  i, o = 99, "201"
  T(_t, S(i), S(convertToBase7(i)), S(o))

  i, o = -7, "-10"
  T(_t, S(i), S(convertToBase7(i)), S(o))
}

