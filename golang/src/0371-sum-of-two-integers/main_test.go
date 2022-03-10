package main
import "testing"
import . "main/pkg/testing_helper"

func TestGetSum(_t *testing.T) {
  i1, i2, o := 1, 2, 3
  T(_t, S(i1, i2), S(getSum(i1, i2)), S(o))

  i1, i2, o = 2, 4, 6
  T(_t, S(i1, i2), S(getSum(i1, i2)), S(o))
}
