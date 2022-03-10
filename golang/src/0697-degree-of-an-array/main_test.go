package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindShortestSubArray(_t *testing.T) {
  i, o := []int{1,2,2,3,1}, 2
  T(_t, S(i), S(findShortestSubArray(i)), S(o))

  i, o = []int{1,2,2,3,1,4,2}, 6
  T(_t, S(i), S(findShortestSubArray(i)), S(o))
}

func TestFindShortestSubArrayMap(_t *testing.T) {
  i, o := []int{1,2,2,3,1}, 2
  T(_t, S(i), S(findShortestSubArrayMap(i)), S(o))

  i, o = []int{1,2,2,3,1,4,2}, 6
  T(_t, S(i), S(findShortestSubArrayMap(i)), S(o))
}

