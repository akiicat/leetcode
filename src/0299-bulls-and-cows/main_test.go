package main
import "testing"
import . "main/pkg/testing_helper"

func TestGetHint(_t *testing.T) {
  i1, i2, o := "1807", "7810", "1A3B"
  T(_t, S(i1, i2), S(getHint(i1, i2)), S(o))

  i1, i2, o = "1123", "0111", "1A1B"
  T(_t, S(i1, i2), S(getHint(i1, i2)), S(o))

  i1, i2, o = "1", "0", "0A0B"
  T(_t, S(i1, i2), S(getHint(i1, i2)), S(o))
}
