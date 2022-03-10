package main
import "testing"
import . "main/pkg/testing_helper"
import "strings"

func TestLongestPalindrome(_t *testing.T) {
  i, o := "babad", "bab"
  T(_t, S(i), S(Valid(i, longestPalindrome(i), o)), S(o))

  i, o = "cbbd", "bb"
  T(_t, S(i), S(Valid(i, longestPalindrome(i), o)), S(o))

  i, o = "a", "a"
  T(_t, S(i), S(Valid(i, longestPalindrome(i), o)), S(o))
}

func TestLongestPalindromeExpandAroundCenter(_t *testing.T) {
  i, o := "babad", "bab"
  T(_t, S(i), S(Valid(i, longestPalindromeExpandAroundCenter(i), o)), S(o))

  i, o = "cbbd", "bb"
  T(_t, S(i), S(Valid(i, longestPalindromeExpandAroundCenter(i), o)), S(o))

  i, o = "a", "a"
  T(_t, S(i), S(Valid(i, longestPalindromeExpandAroundCenter(i), o)), S(o))
}

func TestLongestPalindromeBruteForce(_t *testing.T) {
  i, o := "babad", "bab"
  T(_t, S(i), S(Valid(i, longestPalindromeBruteForce(i), o)), S(o))

  i, o = "cbbd", "bb"
  T(_t, S(i), S(Valid(i, longestPalindromeBruteForce(i), o)), S(o))

  i, o = "a", "a"
  T(_t, S(i), S(Valid(i, longestPalindromeBruteForce(i), o)), S(o))
}

func Valid(s, i, o string) string {

  if !strings.Contains(s, i) {
    return i
  }

  if len(i) != len(o) {
    return i
  }

  left, right := 0, len(i) - 1

  for left < right {
    if i[left] != i[right] {
      return i
    }

    left++
    right--
  }

  return o
}

