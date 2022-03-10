package main
import "testing"
import . "main/pkg/testing_helper"

func TestUniqueOccurrences(t *testing.T) {
  i, o := []int{1,2,2,1,1,3}, true
  T(t, S(i), S(uniqueOccurrences(i)), S(o))

  i, o = []int{1,2}, false
  T(t, S(i), S(uniqueOccurrences(i)), S(o))

  i, o = []int{-3,0,1,-3,1,1,1,-3,10,0}, true
  T(t, S(i), S(uniqueOccurrences(i)), S(o))
}

