package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumberOfSteps(_t *testing.T) {
  i, o := 14, 6
  T(_t, S(i), S(numberOfSteps(i)), S(o))

  i, o = 8, 4
  T(_t, S(i), S(numberOfSteps(i)), S(o))

  i, o = 123, 12
  T(_t, S(i), S(numberOfSteps(i)), S(o))
}

