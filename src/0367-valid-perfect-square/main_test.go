package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsPerfectSquare(_t *testing.T) {
  i, o := 16, true
  T(_t, S(i), S(isPerfectSquare(i)), S(o))

  i, o = 14, false
  T(_t, S(i), S(isPerfectSquare(i)), S(o))
}

func TestIsPerfectSquareMath(_t *testing.T) {
  i, o := 16, true
  T(_t, S(i), S(isPerfectSquareMath(i)), S(o))

  i, o = 14, false
  T(_t, S(i), S(isPerfectSquareMath(i)), S(o))
}
