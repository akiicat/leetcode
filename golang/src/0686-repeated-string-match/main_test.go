package main
import "testing"
import . "main/pkg/testing_helper"

func TestRepeatedStringMatch(_t *testing.T) {
  i1, i2, o := "abcd", "cdabcdab", 3
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "abcd", "cdab", 2
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "a", "aa", 2
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "aa", "aaa", 2
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "abc", "cabcabca", 4
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "ab", "bababababa", 6
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "abcabcabcabc", "abac", -1
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))

  i1, i2, o = "aa", "a", 1
  T(_t, S(i1, i2), S(repeatedStringMatch(i1, i2)), S(o))
}

