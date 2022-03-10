package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinTimeToVisitAllPoints(t *testing.T) {
  i, o := [][]int{[]int{1,1}, []int{3,4}, []int{-1,0}}, 7
  T(t, S(i), S(minTimeToVisitAllPoints(i)), S(o))

  i, o = [][]int{[]int{3,2}, []int{-2,2}}, 5
  T(t, S(i), S(minTimeToVisitAllPoints(i)), S(o))
}

