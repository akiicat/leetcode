package main
import "testing"
import . "main/pkg/testing_helper"

func TestConvert(_t *testing.T) {
  i, n, o := "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"
  T(_t, S(i, n), S(convert(i, n)), S(o))

  i, n, o = "PAYPALISHIRING", 4, "PINALSIGYAHRPI"
  T(_t, S(i, n), S(convert(i, n)), S(o))

  i, n, o = "A", 1, "A"
  T(_t, S(i, n), S(convert(i, n)), S(o))
}

