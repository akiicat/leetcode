package main
import "testing"
import . "main/pkg/testing_helper"

func TestBalancedStringSplit(t *testing.T) {
  i, o := "RLRRLLRLRL", 4
  T(t, S(i), S(balancedStringSplit(i)), S(o))

  i, o = "RLLLLRRRLR", 3
  T(t, S(i), S(balancedStringSplit(i)), S(o))

  i, o = "LLLLRRRR", 1
  T(t, S(i), S(balancedStringSplit(i)), S(o))

  i, o = "RLRRRLLRLL", 2
  T(t, S(i), S(balancedStringSplit(i)), S(o))
}

