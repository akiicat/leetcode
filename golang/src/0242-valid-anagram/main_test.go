package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsAnagram(_t *testing.T) {
  s, t, o := "anagram", "nagaram", true
  T(_t, S(s, t), S(isAnagram(s, t)), S(o))

  s, t, o = "rat", "car", false
  T(_t, S(s, t), S(isAnagram(s, t)), S(o))

  s, t, o = "aa", "bb", false
  T(_t, S(s, t), S(isAnagram(s, t)), S(o))

  s, t, o = "ac", "bb", false
  T(_t, S(s, t), S(isAnagram(s, t)), S(o))

  s, t, o = "aacc", "bbbb", false
  T(_t, S(s, t), S(isAnagram(s, t)), S(o))
}
