package main
import "testing"
import . "main/pkg/testing_helper"

func TestMoveZeroes(_t *testing.T) {
  s, t, o := "abba", "dog cat cat dog", true
  T(_t, S(s, t), S(wordPattern(s, t)), S(o))

  s, t, o = "abba", "dog cat cat fish", false
  T(_t, S(s, t), S(wordPattern(s, t)), S(o))

  s, t, o = "aaaa", "dog cat cat dog", false
  T(_t, S(s, t), S(wordPattern(s, t)), S(o))

  s, t, o = "abba", "dog dog dog dog", false
  T(_t, S(s, t), S(wordPattern(s, t)), S(o))

  s, t, o = "aaa", "dog dog dog dog", false
  T(_t, S(s, t), S(wordPattern(s, t)), S(o))
}
