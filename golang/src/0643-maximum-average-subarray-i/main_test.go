package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindMaxAverage(_t *testing.T) {
  i, k, o := []int{1,12,-5,-6,50,3}, 4, 12.75
  T(_t, S(i, k), S(findMaxAverage(i, k)), S(o))

  i, k, o = []int{5}, 1, 5.00
  T(_t, S(i, k), S(findMaxAverage(i, k)), S(o))
}

