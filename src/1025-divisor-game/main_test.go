package main
import "testing"
import . "main/pkg/testing_helper"

func TestDivisorGame(t *testing.T) {
  i, o := 2, true
  T(t, S(i), S(divisorGame(i)), S(o))

  i, o = 3, false
  T(t, S(i), S(divisorGame(i)), S(o))

  i, o = 6, true
  T(t, S(i), S(divisorGame(i)), S(o))
}

