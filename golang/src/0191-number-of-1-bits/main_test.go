package main
import "testing"
import . "main/pkg/testing_helper"

func TestHammingWeight(_t *testing.T) {
  i, o := uint32(11), 3
  T(_t, S(i), S(hammingWeight(i)), S(o))

  i, o = uint32(128), 1
  T(_t, S(i), S(hammingWeight(i)), S(o))

  i, o = uint32(4294967293), 31
  T(_t, S(i), S(hammingWeight(i)), S(o))
}

func TestHammingWeightForLoop(_t *testing.T) {
  i, o := uint32(11), 3
  T(_t, S(i), S(hammingWeightForLoop(i)), S(o))

  i, o = uint32(128), 1
  T(_t, S(i), S(hammingWeightForLoop(i)), S(o))

  i, o = uint32(4294967293), 31
  T(_t, S(i), S(hammingWeightForLoop(i)), S(o))
}

func TestHammingWeightAlgorithm(_t *testing.T) {
  i, o := uint32(11), 3
  T(_t, S(i), S(hammingWeightAlgorithm(i)), S(o))

  i, o = uint32(128), 1
  T(_t, S(i), S(hammingWeightAlgorithm(i)), S(o))

  i, o = uint32(4294967293), 31
  T(_t, S(i), S(hammingWeightAlgorithm(i)), S(o))
}
