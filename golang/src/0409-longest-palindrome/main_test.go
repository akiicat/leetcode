package main
import "testing"
import . "main/pkg/testing_helper"

func TestLongestPalindrome(_t *testing.T) {
  i, o := "abccccdd", 7
  T(_t, S(i), S(longestPalindrome(i)), S(o))
}
