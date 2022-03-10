package main
import "testing"
import . "main/pkg/testing_helper"

func TestDetectCapitalUse(_t *testing.T) {
  i, o := "USA", true
  T(_t, S(i), S(detectCapitalUse(i)), S(o))

  i, o = "FlaG", false
  T(_t, S(i), S(detectCapitalUse(i)), S(o))

  i, o = "Flag", true
  T(_t, S(i), S(detectCapitalUse(i)), S(o))
}

func TestDetectCapitalUseString(_t *testing.T) {
  i, o := "USA", true
  T(_t, S(i), S(detectCapitalUseString(i)), S(o))

  i, o = "FlaG", false
  T(_t, S(i), S(detectCapitalUseString(i)), S(o))

  i, o = "Flag", true
  T(_t, S(i), S(detectCapitalUseString(i)), S(o))
}
