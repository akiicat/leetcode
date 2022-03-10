package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumPairsDivisibleBy60(t *testing.T) {
  i, o := []int{30,20,150,100,40}, 3
  T(t, S(i), S(numPairsDivisibleBy60(i)), S(o))

  i, o = []int{60,60,60}, 3
  T(t, S(i), S(numPairsDivisibleBy60(i)), S(o))

  i, o = []int{15,63,451,213,37,209,343,319}, 1
  T(t, S(i), S(numPairsDivisibleBy60(i)), S(o))
}

