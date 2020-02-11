package main
import "testing"
import . "main/pkg/testing_helper"

func TestValidPalindrome(_t *testing.T) {
  i, o := "aba", true
  T(_t, S(i), S(validPalindrome(i)), S(o))

  i, o = "abca", true
  T(_t, S(i), S(validPalindrome(i)), S(o))

  i, o = "abbca", true
  T(_t, S(i), S(validPalindrome(i)), S(o))

  i, o = "abcbca", true
  T(_t, S(i), S(validPalindrome(i)), S(o))

  i, o = "abc", false
  T(_t, S(i), S(validPalindrome(i)), S(o))
}
