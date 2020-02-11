package main
import "testing"
import . "main/pkg/testing_helper"

func TestLongestWord(_t *testing.T) {
  i, o := []string{"w","wo","wor","worl", "world"}, "world"
  T(_t, S(i), S(longestWord(i)), S(o))

  i, o = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"
  T(_t, S(i), S(longestWord(i)), S(o))

  i, o = []string{"a", "b", "cc"}, "a"
  T(_t, S(i), S(longestWord(i)), S(o))

  i, o = []string{"w","wo","wox","worl", "world"}, "wox"
  T(_t, S(i), S(longestWord(i)), S(o))
}

func TestLongestWordSort(_t *testing.T) {
  i, o := []string{"w","wo","wor","worl", "world"}, "world"
  T(_t, S(i), S(longestWordSort(i)), S(o))

  i, o = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"
  T(_t, S(i), S(longestWordSort(i)), S(o))

  i, o = []string{"a", "b", "cc"}, "a"
  T(_t, S(i), S(longestWordSort(i)), S(o))

  i, o = []string{"w","wo","wox","worl", "world"}, "wox"
  T(_t, S(i), S(longestWordSort(i)), S(o))
}

