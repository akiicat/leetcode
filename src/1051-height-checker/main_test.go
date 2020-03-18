package main
import "testing"
import . "main/pkg/testing_helper"

func TestHeightChecker(_t *testing.T) {
  i, o := []int{1,1,4,2,1,3}, 3
  T(_t, S(i), S(heightChecker(i)), S(o))

  i, o = []int{5,1,2,3,4}, 5
  T(_t, S(i), S(heightChecker(i)), S(o))

  i, o = []int{1,2,3,4,5}, 0
  T(_t, S(i), S(heightChecker(i)), S(o))
}

