package main
import "testing"
import . "main/pkg/testing_helper"

func TestBuddyStrings(_t *testing.T) {
  a, b, o := "ab", "ba", true
  T(_t, S(a, b), S(buddyStrings(a, b)), S(o))

  a, b, o = "ab", "ab", false
  T(_t, S(a, b), S(buddyStrings(a, b)), S(o))

  a, b, o = "aa", "aa", true
  T(_t, S(a, b), S(buddyStrings(a, b)), S(o))

  a, b, o = "abcd", "badc", false
  T(_t, S(a, b), S(buddyStrings(a, b)), S(o))

  a, b, o = "aaaaaaabc", "aaaaaaacb", true
  T(_t, S(a, b), S(buddyStrings(a, b)), S(o))
}

func TestBuddyStringsIntuition(_t *testing.T) {
  a, b, o := "ab", "ba", true
  T(_t, S(a, b), S(buddyStringsIntuition(a, b)), S(o))

  a, b, o = "ab", "ab", false
  T(_t, S(a, b), S(buddyStringsIntuition(a, b)), S(o))

  a, b, o = "aa", "aa", true
  T(_t, S(a, b), S(buddyStringsIntuition(a, b)), S(o))

  a, b, o = "abcd", "badc", false
  T(_t, S(a, b), S(buddyStringsIntuition(a, b)), S(o))

  a, b, o = "aaaaaaabc", "aaaaaaacb", true
  T(_t, S(a, b), S(buddyStringsIntuition(a, b)), S(o))
}

