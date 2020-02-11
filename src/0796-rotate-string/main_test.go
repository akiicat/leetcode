package main
import "testing"
import . "main/pkg/testing_helper"

func TestRotateString(_t *testing.T) {
  a, b, o := "abcde", "cdeab", true
  T(_t, S(a, b), S(rotateString(a, b)), S(o))

  a, b, o = "abcde", "abced", false
  T(_t, S(a, b), S(rotateString(a, b)), S(o))
}

func TestRotateStringSimpleCheck(_t *testing.T) {
  a, b, o := "abcde", "cdeab", true
  T(_t, S(a, b), S(rotateStringSimpleCheck(a, b)), S(o))

  a, b, o = "abcde", "abced", false
  T(_t, S(a, b), S(rotateStringSimpleCheck(a, b)), S(o))
}

