package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsPalindrome(_t *testing.T) {
  i, o := "A man, a plan, a canal: Panama", true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = "race a car", false
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = "OP", false
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = "0P", false
  T(_t, S(i), S(isPalindrome(i)), S(o))
}
