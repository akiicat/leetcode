package main
import "testing"
import . "main/pkg/testing_helper"

func TestAddStrings(_t *testing.T) {
  i1, i2, o := "0", "0", "0"
  T(_t, S(i1, i2), S(addStrings(i1, i2)), S(o))

  i1, i2, o = "10", "90", "100"
  T(_t, S(i1, i2), S(addStrings(i1, i2)), S(o))

  i1, i2, o = "9", "99", "108"
  T(_t, S(i1, i2), S(addStrings(i1, i2)), S(o))

  i1, i2, o = "5432", "54", "5486"
  T(_t, S(i1, i2), S(addStrings(i1, i2)), S(o))
}

