package main
import "testing"
import . "main/pkg/testing_helper"

func TestLargestTimeFromDigits(_t *testing.T) {
  i, o := []int{1,2,3,4}, "23:41"
  T(_t, S(i), S(largestTimeFromDigits(i)), S(o))

  i, o = []int{5,5,5,5}, ""
  T(_t, S(i), S(largestTimeFromDigits(i)), S(o))

  i, o = []int{2,0,6,6}, "06:26"
  T(_t, S(i), S(largestTimeFromDigits(i)), S(o))
}


func TestLargestTimeFromDigitsBruteForce(_t *testing.T) {
  i, o := []int{1,2,3,4}, "23:41"
  T(_t, S(i), S(largestTimeFromDigitsBruteForce(i)), S(o))

  i, o = []int{5,5,5,5}, ""
  T(_t, S(i), S(largestTimeFromDigitsBruteForce(i)), S(o))

  i, o = []int{2,0,6,6}, "06:26"
  T(_t, S(i), S(largestTimeFromDigitsBruteForce(i)), S(o))
}
