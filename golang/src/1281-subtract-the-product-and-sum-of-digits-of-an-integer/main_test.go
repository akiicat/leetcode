package main
import "testing"
import . "main/pkg/testing_helper"

func TestSubtractProductAndSum(t *testing.T) {
  i, o := 234, 15
  T(t, S(i), S(subtractProductAndSum(i)), S(o))

  i, o = 4421, 21
  T(t, S(i), S(subtractProductAndSum(i)), S(o))

  i, o = 111, -2
  T(t, S(i), S(subtractProductAndSum(i)), S(o))
}

