package main
import "testing"
import . "main/pkg/testing_helper"

func TestHammingDistance(_t *testing.T) {
  i1, i2, o := 1, 4, 2
  T(_t, S(i1, i2), S(hammingDistance(i1, i2)), S(o))

  i1, i2, o = 93, 73, 2
  T(_t, S(i1, i2), S(hammingDistance(i1, i2)), S(o))
}

