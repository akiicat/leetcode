package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsIsomorphic(_t *testing.T) {
  s, t, o := "egg", "add", true
  T(_t, S(s, t), S(isIsomorphic(s, t)), S(o))

  s, t, o = "foo", "bar", false
  T(_t, S(s, t), S(isIsomorphic(s, t)), S(o))

  s, t, o = "paper", "title", true
  T(_t, S(s, t), S(isIsomorphic(s, t)), S(o))

  s, t, o = "ab", "aa", false
  T(_t, S(s, t), S(isIsomorphic(s, t)), S(o))
}
