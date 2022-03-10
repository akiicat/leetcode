package main
import "testing"
import . "main/pkg/testing_helper"

func TestValidMountainArray(_t *testing.T) {
  i, o := []int{}, false
  T(_t, S(i), S(validMountainArray(i)), S(o))

  i, o = []int{2,1}, false
  T(_t, S(i), S(validMountainArray(i)), S(o))

  i, o = []int{3,5,5}, false
  T(_t, S(i), S(validMountainArray(i)), S(o))

  i, o = []int{0,3,2,1}, true
  T(_t, S(i), S(validMountainArray(i)), S(o))

  i, o = []int{0,2,3,4,5,2,1,0}, true
  T(_t, S(i), S(validMountainArray(i)), S(o))

  i, o = []int{0,2,3,3,5,2,1,0}, false
  T(_t, S(i), S(validMountainArray(i)), S(o))
}

