package main
import "testing"
import . "main/pkg/testing_helper"

func TestRemoveDuplicates(_t *testing.T) {
  i, o := "abbaca", "ca"
  T(_t, S(i), S(removeDuplicates(i)), S(o))

  i, o = "aaaaaaaaa", "a"
  T(_t, S(i), S(removeDuplicates(i)), S(o))

  i, o = "aaaaaaaa", ""
  T(_t, S(i), S(removeDuplicates(i)), S(o))
}

func TestRemoveDuplicatesString(_t *testing.T) {
  i, o := "abbaca", "ca"
  T(_t, S(i), S(removeDuplicatesString(i)), S(o))

  i, o = "aaaaaaaaa", "a"
  T(_t, S(i), S(removeDuplicatesString(i)), S(o))

  i, o = "aaaaaaaa", ""
  T(_t, S(i), S(removeDuplicatesString(i)), S(o))
}

