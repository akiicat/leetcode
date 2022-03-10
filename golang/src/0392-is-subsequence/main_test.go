package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindTheDifference(_t *testing.T) {
  s, t, o := "abc", "ahbgdc", true
  T(_t, S(s, t), S(isSubsequence(s, t)), S(o))

  s, t, o = "axc", "ahbgdc", false
  T(_t, S(s, t), S(isSubsequence(s, t)), S(o))

  s, t, o = "aec", "abcde", false
  T(_t, S(s, t), S(isSubsequence(s, t)), S(o))
}

func TestIsSubsequenceEarlyStop(_t *testing.T) {
  s, t, o := "abc", "ahbgdc", true
  T(_t, S(s, t), S(isSubsequenceEarlyStop(s, t)), S(o))

  s, t, o = "axc", "ahbgdc", false
  T(_t, S(s, t), S(isSubsequenceEarlyStop(s, t)), S(o))

  s, t, o = "aec", "abcde", false
  T(_t, S(s, t), S(isSubsequenceEarlyStop(s, t)), S(o))
}
