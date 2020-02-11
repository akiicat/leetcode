package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountPrimeSetBits(_t *testing.T) {
  l, r, o := 6, 10, 4
  T(_t, S(l, r), S(countPrimeSetBits(l, r)), S(o))

  l, r, o = 10, 15, 5
  T(_t, S(l, r), S(countPrimeSetBits(l, r)), S(o))

  l, r, o = 842, 888, 23
  T(_t, S(l, r), S(countPrimeSetBits(l, r)), S(o))
}


