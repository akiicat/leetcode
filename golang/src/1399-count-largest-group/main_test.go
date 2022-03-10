package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountLargestGroup(_t *testing.T) {
  i, o := 13, 4
  T(_t, S(i), S(countLargestGroup(i)), S(o))

  i, o = 2, 2
  T(_t, S(i), S(countLargestGroup(i)), S(o))

  i, o = 15, 6
  T(_t, S(i), S(countLargestGroup(i)), S(o))

  i, o = 24, 5
  T(_t, S(i), S(countLargestGroup(i)), S(o))
}

