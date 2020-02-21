package main
import "testing"
import . "main/pkg/testing_helper"

func TestMyPow(t *testing.T) {
  x, n, o := 2.00000, 10, 1024.00000
  T(t, S(x, n), S(myPow(x, n)), S(o))

  x, n, o = 2.10000, 3, 9.26100
  T(t, S(x, n), S(myPow(x, n)), S(o))

  x, n, o = 2.00000, -2, 0.25000
  T(t, S(x, n), S(myPow(x, n)), S(o))
}

func TestMyPowLoop(t *testing.T) {
  x, n, o := 2.00000, 10, 1024.00000
  T(t, S(x, n), S(myPowLoop(x, n)), S(o))

  x, n, o = 2.10000, 3, 9.26100
  T(t, S(x, n), S(myPowLoop(x, n)), S(o))

  x, n, o = 2.00000, -2, 0.25000
  T(t, S(x, n), S(myPowLoop(x, n)), S(o))
}

