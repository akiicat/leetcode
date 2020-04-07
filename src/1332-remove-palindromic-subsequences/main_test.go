package main
import "testing"
import . "main/pkg/testing_helper"

func TestRemovePalindromeSub(t *testing.T) {
  i, o := "ababa", 1
  T(t, S(i), S(removePalindromeSub(i)), S(o))

  i, o = "abb", 2
  T(t, S(i), S(removePalindromeSub(i)), S(o))

  i, o = "baabb", 2
  T(t, S(i), S(removePalindromeSub(i)), S(o))

  i, o = "abbaaaab", 2
  T(t, S(i), S(removePalindromeSub(i)), S(o))

  i, o = "", 0
  T(t, S(i), S(removePalindromeSub(i)), S(o))
}

