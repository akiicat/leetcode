package main
import "testing"
import . "main/pkg/testing_helper"

func TestBinaryGap(_t *testing.T) {
  i, o := 22, 2
  T(_t, S(i), S(binaryGap(i)), S(o))

  i, o = 5, 2
  T(_t, S(i), S(binaryGap(i)), S(o))

  i, o = 6, 1
  T(_t, S(i), S(binaryGap(i)), S(o))

  i, o = 8, 0
  T(_t, S(i), S(binaryGap(i)), S(o))
}

