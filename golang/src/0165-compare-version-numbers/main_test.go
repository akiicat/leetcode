package main
import "testing"
import . "main/pkg/testing_helper"

func TestCompareVersion(_t *testing.T) {
  i1, i2, o := "0.1", "1.1", -1
  T(_t, S(i1, i2), S(compareVersion(i1, i2)), S(o))

  i1, i2, o = "1.0.1", "1", 1
  T(_t, S(i1, i2), S(compareVersion(i1, i2)), S(o))

  i1, i2, o = "7.5.2.4", "7.5.3", -1
  T(_t, S(i1, i2), S(compareVersion(i1, i2)), S(o))

  i1, i2, o = "1.01", "1.001", 0
  T(_t, S(i1, i2), S(compareVersion(i1, i2)), S(o))

  i1, i2, o = "1.0", "1.0.0", 0
  T(_t, S(i1, i2), S(compareVersion(i1, i2)), S(o))
}

