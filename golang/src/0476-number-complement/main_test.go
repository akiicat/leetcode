package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindComplement(_t *testing.T) {
  i, o := 5, 2
  T(_t, S(i), S(findComplement(i)), S(o))

  i, o = 6, 1
  T(_t, S(i), S(findComplement(i)), S(o))

  i, o = 1, 0
  T(_t, S(i), S(findComplement(i)), S(o))
}

func TestFindComplementShiftLeft(_t *testing.T) {
  i, o := 5, 2
  T(_t, S(i), S(findComplementShiftLeft(i)), S(o))

  i, o = 6, 1
  T(_t, S(i), S(findComplementShiftLeft(i)), S(o))

  i, o = 1, 0
  T(_t, S(i), S(findComplementShiftLeft(i)), S(o))
}
