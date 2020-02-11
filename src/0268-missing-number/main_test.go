package main
import "testing"
import . "main/pkg/testing_helper"

func TestMissingNumber(_t *testing.T) {
  i, o := []int{3,0,1}, 2
  T(_t, S(i), S(missingNumber(i)), S(o))

  i, o = []int{9,6,4,2,3,5,7,0,1}, 8
  T(_t, S(i), S(missingNumber(i)), S(o))
}

func TestMissingNumberBitManipulation(_t *testing.T) {
  i, o := []int{3,0,1}, 2
  T(_t, S(i), S(missingNumberBitManipulation(i)), S(o))

  i, o = []int{9,6,4,2,3,5,7,0,1}, 8
  T(_t, S(i), S(missingNumberBitManipulation(i)), S(o))
}
