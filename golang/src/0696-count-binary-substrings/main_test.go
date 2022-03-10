package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountBinarySubstrings(_t *testing.T) {
  i, o := "00110011", 6
  T(_t, S(i), S(countBinarySubstrings(i)), S(o))

  i, o = "10101", 4
  T(_t, S(i), S(countBinarySubstrings(i)), S(o))
}

func TestCountBinarySubstringsArray(_t *testing.T) {
  i, o := "00110011", 6
  T(_t, S(i), S(countBinarySubstringsArray(i)), S(o))

  i, o = "10101", 4
  T(_t, S(i), S(countBinarySubstringsArray(i)), S(o))
}

