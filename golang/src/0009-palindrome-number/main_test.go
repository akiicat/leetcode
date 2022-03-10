package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsPalindrome(_t *testing.T) {
  i, o := 0, true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = 1, true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = 121, true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = -121, false
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = 313, true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = 1122, false
  T(_t, S(i), S(isPalindrome(i)), S(o))
}

