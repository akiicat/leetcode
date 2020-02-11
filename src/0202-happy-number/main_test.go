package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsHappy(_t *testing.T) {
  i, o := 19, true
  T(_t, S(i), S(isHappy(i)), S(o))

  i, o = 2, false
  T(_t, S(i), S(isHappy(i)), S(o))
}

func TestIsHappyHash(_t *testing.T) {
  i, o := 19, true
  T(_t, S(i), S(isHappyHash(i)), S(o))

  i, o = 2, false
  T(_t, S(i), S(isHappyHash(i)), S(o))
}

