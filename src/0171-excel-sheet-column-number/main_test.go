package main
import "testing"
import . "main/pkg/testing_helper"

func TestTitleToNumber(_t *testing.T) {
  i, o := "A", 1
  T(_t, S(i), S(titleToNumber(i)), S(o))

  i, o = "Z", 26
  T(_t, S(i), S(titleToNumber(i)), S(o))

  i, o = "AB", 28
  T(_t, S(i), S(titleToNumber(i)), S(o))

  i, o = "AZ", 52
  T(_t, S(i), S(titleToNumber(i)), S(o))

  i, o = "ZY", 701
  T(_t, S(i), S(titleToNumber(i)), S(o))

  i, o = "AAA", 703
  T(_t, S(i), S(titleToNumber(i)), S(o))
}

