package main
import "testing"
import . "main/pkg/testing_helper"

func TestLongestPalindrome(_t *testing.T) {
  i, o := "abc", 3
  T(_t, S(i), S(countSubstrings(i)), S(o))

  i, o = "babad", 7
  T(_t, S(i), S(countSubstrings(i)), S(o))

  i, o = "aaa", 6
  T(_t, S(i), S(countSubstrings(i)), S(o))

  i, o = "a", 1
  T(_t, S(i), S(countSubstrings(i)), S(o))

  i, o = "", 0
  T(_t, S(i), S(countSubstrings(i)), S(o))
}

func TestLongestPalindromeExpandAroundCenter(_t *testing.T) {
  i, o := "abc", 3
  T(_t, S(i), S(countSubstringsExpandAroundCenter(i)), S(o))

  i, o = "babad", 7
  T(_t, S(i), S(countSubstringsExpandAroundCenter(i)), S(o))

  i, o = "aaa", 6
  T(_t, S(i), S(countSubstringsExpandAroundCenter(i)), S(o))

  i, o = "a", 1
  T(_t, S(i), S(countSubstringsExpandAroundCenter(i)), S(o))

  i, o = "", 0
  T(_t, S(i), S(countSubstringsExpandAroundCenter(i)), S(o))
}

