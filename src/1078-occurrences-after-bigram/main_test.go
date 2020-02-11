package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindOcurrences(t *testing.T) {
  i, first, second, o := "alice is a good girl she is a good student", "a", "good", []string{"girl", "student"}
  T(t, S(i, first, second), S(findOcurrences(i, first, second)), S(o))

  i, first, second, o = "we will we will rock you", "we", "will", []string{"we", "rock"}
  T(t, S(i, first, second), S(findOcurrences(i, first, second)), S(o))

  i, first, second, o = "a a a a a a alice is a good girl she is a good student", "a", "a", []string{"a","a","a","a","alice"}
  T(t, S(i, first, second), S(findOcurrences(i, first, second)), S(o))
}


