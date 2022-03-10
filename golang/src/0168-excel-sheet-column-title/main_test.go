package main
import "testing"
import . "main/pkg/testing_helper"

func TestConvertToTitle(_t *testing.T) {
  i, o := 1, "A"
  T(_t, S(i), S(convertToTitle(i)), S(o))

  i, o = 26, "Z"
  T(_t, S(i), S(convertToTitle(i)), S(o))

  i, o = 28, "AB"
  T(_t, S(i), S(convertToTitle(i)), S(o))

  i, o = 52, "AZ"
  T(_t, S(i), S(convertToTitle(i)), S(o))

  i, o = 701, "ZY"
  T(_t, S(i), S(convertToTitle(i)), S(o))

  i, o = 703, "AAA"
  T(_t, S(i), S(convertToTitle(i)), S(o))
}

