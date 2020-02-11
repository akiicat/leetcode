package main
import "testing"
import . "main/pkg/testing_helper"

func TestHasGroupsSizeX(t *testing.T) {
  i, o := []int{1,2,3,4,4,3,2,1}, true
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{1,1,1,2,2,2,3,3}, false
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{1}, false
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{1,1}, true
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{0,0,0,1,1,1,2,2,2}, true
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{1,1,2,2,2,2}, true
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{1,1,1,1,2,2,2,2,2,2}, true
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))

  i, o = []int{1,1,1,1,1,1,2,2,2,2,2,2,2,2,2,3,3,3,3,3,3,3,3}, false
  T(t, S(i), S(hasGroupsSizeX(i)), S(o))
}

