package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumJewelsInStones(_t *testing.T) {
  j, s, o := "aA", "aAAbbbb", 3
  T(_t, S(j, s), S(numJewelsInStones(j, s)), S(o))

  j, s, o = "z", "ZZ", 0
  T(_t, S(j, s), S(numJewelsInStones(j, s)), S(o))
}

