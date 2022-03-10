package main
import "testing"
import . "main/pkg/testing_helper"

func TestPowerfulIntegers(t *testing.T) {
  i, o := 5, 2
  T(t, S(i), S(bitwiseComplement(i)), S(o))

  i, o = 7, 0
  T(t, S(i), S(bitwiseComplement(i)), S(o))

  i, o = 10, 5
  T(t, S(i), S(bitwiseComplement(i)), S(o))

  i, o = 0, 1
  T(t, S(i), S(bitwiseComplement(i)), S(o))
}

