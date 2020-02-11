package main
import "testing"
import . "main/pkg/testing_helper"

func TestRepeatedSubstringPattern(_t *testing.T) {
  i, o := "abab", true
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "aba", false
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "abcabcabcabc", true
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "ababba", false
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "aabaaba", false
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "bb", true
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "ababababab", true
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "abaababaab", true
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))

  i, o = "ab", false
  T(_t, S(i), S(repeatedSubstringPattern(i)), S(o))
}

func TestRepeatedSubstringPatternRepeat(_t *testing.T) {
  i, o := "abab", true
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "aba", false
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "abcabcabcabc", true
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "ababba", false
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "aabaaba", false
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "bb", true
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "ababababab", true
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "abaababaab", true
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))

  i, o = "ab", false
  T(_t, S(i), S(repeatedSubstringPatternRepeat(i)), S(o))
}
