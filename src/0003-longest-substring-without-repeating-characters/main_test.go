package main
import "testing"
import . "main/pkg/testing_helper"

func TestLengthOfLongestSubstring(_t *testing.T) {
  i, o := "abcabcbb", 3
  T(_t, S(i), S(lengthOfLongestSubstring(i)), S(o))

  i, o = "bbbbb", 1
  T(_t, S(i), S(lengthOfLongestSubstring(i)), S(o))

  i, o = "pwwkew", 3
  T(_t, S(i), S(lengthOfLongestSubstring(i)), S(o))

  i, o = " ", 1
  T(_t, S(i), S(lengthOfLongestSubstring(i)), S(o))
}

