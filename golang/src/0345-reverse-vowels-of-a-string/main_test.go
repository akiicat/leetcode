package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverseVowels(_t *testing.T) {
  i, o := "hello", "holle"
  T(_t, S(i), S(reverseVowels(i)), S(o))

  i, o = "leetcode", "leotcede"
  T(_t, S(i), S(reverseVowels(i)), S(o))

  i, o = "aA", "Aa"
  T(_t, S(i), S(reverseVowels(i)), S(o))
}

