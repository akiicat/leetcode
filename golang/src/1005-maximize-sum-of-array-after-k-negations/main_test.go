package main
import "testing"
import . "main/pkg/testing_helper"

func TestLargestSumAfterKNegations(t *testing.T) {
  i, k, o := []int{4,2,3}, 1, 5
  T(t, S(i, k), S(largestSumAfterKNegations(i, k)), S(o))

  i, k, o = []int{3,-1,0,2}, 3, 6
  T(t, S(i, k), S(largestSumAfterKNegations(i, k)), S(o))

  i, k, o = []int{2,-3,-1,5,-4}, 2, 13
  T(t, S(i, k), S(largestSumAfterKNegations(i, k)), S(o))
}


func TestLargestSumAfterKNegationsSort(t *testing.T) {
  i, k, o := []int{4,2,3}, 1, 5
  T(t, S(i, k), S(largestSumAfterKNegationsSort(i, k)), S(o))

  i, k, o = []int{3,-1,0,2}, 3, 6
  T(t, S(i, k), S(largestSumAfterKNegationsSort(i, k)), S(o))

  i, k, o = []int{2,-3,-1,5,-4}, 2, 13
  T(t, S(i, k), S(largestSumAfterKNegationsSort(i, k)), S(o))
}
